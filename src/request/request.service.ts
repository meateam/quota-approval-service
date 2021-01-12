import * as grpc from "grpc";
import AdminManager from "../admin/admin.service";
import config from "../config";
import { loadSync } from "@grpc/proto-loader";
import RequestRepository from "./request.repository";
import { UnauthorizedError, ServerError, ClientError } from "../utils/error";
import {
    RequestStatus,
    IQuotaApprovalRequest,
    GetRequestsQuery,
} from "./request.interface";
import { status } from "../utils/grpc/status";

const PROTO_PATH = `${__dirname}/../../proto/quota/quota.proto`;
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
        const request = await RequestRepository.getRequestById(requestId);

        if (!request) {
            throw new ClientError(
                `request with id: ${requestId} was not found!`
            );
        }

        return request;
    }

    static async create(
        request: IQuotaApprovalRequest
    ): Promise<IQuotaApprovalRequest> {
        request.status = await this.getInitialRequestStatus(request);
        const requestDocument = await RequestRepository.createRequest(request);

        await this.handleQuotaUpdate(requestDocument);

        return requestDocument;
    }

    static async getAllRequests() {
        return RequestRepository.getRequests({});
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

    static async getApprovableRequests(
        userId: string
    ): Promise<IQuotaApprovalRequest[]> {
        const isAdmin = await AdminManager.isUserAdmin(userId);

        if (!isAdmin) {
            return [];
        }

        const condition = {
            status: RequestStatus.PENDING,
        };

        return RequestRepository.getRequests(condition);
    }

    static async getUserRequests(userId: string) {
        const condition = {
            from: userId,
        };

        return RequestRepository.getRequests(condition);
    }

    static async updateRequest(
        requestId: string,
        modifiedBy: string,
        status: RequestStatus
    ): Promise<IQuotaApprovalRequest> {
        const isAdmin = await AdminManager.isUserAdmin(modifiedBy);

        if (!isAdmin) {
            throw new UnauthorizedError(
                `Operation not permitted for user: ${modifiedBy}`
            );
        }

        const updatedRequest = await RequestRepository.updateRequest(
            requestId,
            {
                status,
                modifiedBy,
            }
        );

        if (!updatedRequest) {
            throw new ClientError(
                `request with id: ${requestId} was not found!`
            );
        }

        await this.handleQuotaUpdate(updatedRequest);

        return updatedRequest;
    }

    /**
     * Returns the initial status for the new request.
     * @param request the new request.
     */
    private static async getInitialRequestStatus(
        request: IQuotaApprovalRequest
    ) {
        if (request.modifiedBy) {
            const isAdmin = await AdminManager.isUserAdmin(request.modifiedBy);

            if (isAdmin) {
                return RequestStatus.APPROVED;
            }
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
    private static async handleQuotaUpdate(
        request: IQuotaApprovalRequest
    ): Promise<void> {
        try {
            if (this.isApproved(request.status) && !config.debugMode) {
                const ownerID: string = request.from;
                const { size } = request;

                const quotaClient = await new quota.QuotaService(
                    config.quotaService.url,
                    grpc.credentials.createInsecure()
                );
                await new Promise((res, rej) => {
                    quotaClient.ChangeQuotaLimit(
                        {
                            ownerID,
                            size,
                        },
                        (err: Error, response: any) => {
                            if (err) {
                                rej(err);
                            }
                            res(response);
                        }
                    );
                });
            }
        } catch (err) {
            const requestId: string = request.id;

            await RequestRepository.updateRequest(requestId, {
                status: RequestStatus.PENDING,
            });

            throw new ServerError(
                `Failed to update Quota service about request change with error: ${JSON.stringify(
                    err.message
                )}`,
                status.INTERNAL
            );
        }
    }
}
