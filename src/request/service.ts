import AdminManager from '../admin/service';
import RequestModel from './model';
import * as grpc from 'grpc';
import { loadSync } from '@grpc/proto-loader';
import { UpdateQuery } from 'mongoose';
import { UnauthorizedError, ServerError, ClientError } from '../utils/error';
import { RequestStatus, IQuotaApprovalRequest, GetRequestsQuery } from './interface';
import config from '../config';

const PROTO_PATH = __dirname + '/../../../proto/quota/quota.proto';
const packageDefinition = loadSync(PROTO_PATH, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
});
const { quota }: any = grpc.loadPackageDefinition(packageDefinition);

export default class RequestService {
    static async getById(requestId: string) {
        const request = await RequestModel.findById(requestId);

        if (!request) {
            throw new ClientError(`request with id: ${requestId} was not found!`);
        }

        return request;
    }

    static async create(request: IQuotaApprovalRequest): Promise<IQuotaApprovalRequest> {
        request.status = await this.getInitialRequestStatus(request);
        const requestDocument = await RequestModel.create(request);

        await this.handleQuotaUpdate(requestDocument);

        return requestDocument;
    }

    static async getAllRequests() {
        return RequestModel.find().exec();
    }

    static async getRequests(query: GetRequestsQuery) {
        const { createdBy, approvableBy } = query;

        if (createdBy) {
            return this.getUserRequests(createdBy);
        }

        if (approvableBy) {
            return this.getApprovableRequests(approvableBy);
        }

        return this.getAllRequests();
    }

    static async getApprovableRequests(userId: string, search?: string): Promise<IQuotaApprovalRequest[]> {
        const isAdmin = await AdminManager.isUserAdmin(userId);

        if (!isAdmin) {
            return [];
        }

        let condition = {
            status: RequestStatus.PENDING,
        };

        if (search) {
            const searchCondition = this.getSearchCondition(search);
            condition = { ...condition, ...searchCondition };
        }

        return RequestModel.find(condition).exec();
    }

    static async getUserRequests(userId: string, search?: string) {
        let condition = {
            from: userId,
        };

        if (search) {
            const searchCondition = this.getSearchCondition(search);
            condition = { ...condition, ...searchCondition };
        }

        return RequestModel.find(condition).exec();
    }

    static async updateRequest(requestId: string, userId: string, status: RequestStatus): Promise<IQuotaApprovalRequest> {
        const isAdmin = await AdminManager.isUserAdmin(userId);

        if (!isAdmin) {
            throw new UnauthorizedError(`Operation not permitted for user: ${userId}`);
        }

        const updateExpression: UpdateQuery<IQuotaApprovalRequest> = {
            $set: {
                status,
            },
        };

        const updatedRequest = await RequestModel.findOneAndUpdate({ id: requestId }, updateExpression, { new: true }).exec();

        if (!updatedRequest) {
            throw new ClientError(`request with id: ${requestId} was not found!`);
        }

        await this.handleQuotaUpdate(updatedRequest);

        return updatedRequest;
    }

    private static getSearchCondition(search: string) {
        return {
            $or: [{ id: { $regex: `^${search}` } }, { from: { $regex: search, $options: 'i' } }],
        };
    }

    /**
     * Returns the initial status for the new request.
     * @param request the new request.
     */
    private static async getInitialRequestStatus(request: IQuotaApprovalRequest) {
        const isAdmin = await AdminManager.isUserAdmin(request.from);

        if (isAdmin) {
            return RequestStatus.APPROVED;
        }

        return RequestStatus.PENDING;
    }

    /**
     * Returns true if the status of the request is approved.
     * @param status the request status
     */
    private static isApproved(status: RequestStatus) {
        return status === RequestStatus.APPROVED;
    }

    /**
     * When the status of a request changes, this function checks if it was approved.
     * If it was approved, it updates the quota in quota-service. If the update fails,
     * the function rolls back the request to its last status.
     * @param request the request that its status was updated.
     */
    private static async handleQuotaUpdate(request: IQuotaApprovalRequest): Promise<void> {
        try {
            if (this.isApproved(request.status)) {
                const ownerID: string = request.from;
                const size: number = request.size;

                const quotaClient = await new quota.QuotaService(config.quotaService.url, grpc.credentials.createInsecure());
                quotaClient.UpdateQuota({ ownerID, size }, (err: Error, _res: any) => {
                    if (err) throw new ServerError();
                });
            }
        } catch (err) {
            const requestId: string = request.id;
            const undoExpression: UpdateQuery<IQuotaApprovalRequest> = {
                $set: {
                    status: request.status,
                },
            };

            await RequestModel.findOneAndUpdate({ id: requestId }, undoExpression, { new: true }).exec();
            throw new ServerError(`Failed to update Quota service about request change with error: ${JSON.stringify(err)}`);
        }
    }
}
