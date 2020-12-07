import { ServerUnaryCall } from 'grpc';
import RequestService from './service';
import { IQuotaApprovalRequest, GetRequestsQuery, RequestStatus } from './interface';

export default class RequestMethods {
    static async createRequest(call: ServerUnaryCall<{ from: string; info: string; size: number; modifiedBy: string; }>): Promise<IQuotaApprovalRequest> {
        const from: string = call.request.from;
        const info: string = call.request.info;
        const size: number = call.request.size;
        const modifiedBy: string = call.request.modifiedBy;
        
        return RequestService.create({ from, info, size, modifiedBy } as IQuotaApprovalRequest);
    }

    static async getRequestById(call: ServerUnaryCall<{ id: string }>): Promise<IQuotaApprovalRequest> {
        const requestId: string = call.request.id;

        return RequestService.getById(requestId);
    }

    static async getRequests(call: ServerUnaryCall<GetRequestsQuery>): Promise<Array<IQuotaApprovalRequest>> {
        const query: GetRequestsQuery = call.request;

        return RequestService.getRequests(query);
    }

    static async updateRequest(call: ServerUnaryCall<{ id: string; modifiedBy: string; status: string }>): Promise<IQuotaApprovalRequest> {
        const id: string = call.request.id;
        const modifiedBy: string = call.request.modifiedBy;
        const status: RequestStatus = call.request.status as RequestStatus;

        return RequestService.updateRequest(id, modifiedBy, status);
    }
}
