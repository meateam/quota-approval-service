import { ServerUnaryCall } from "grpc";
import RequestService from "./request.service";
import {
    IQuotaApprovalRequest,
    GetRequestsQuery,
    RequestStatus,
} from "./request.interface";

export default class RequestMethods {
    static async createRequest(
        call: ServerUnaryCall<{
            from: string;
            info: string;
            size: number;
            modifiedBy: string;
        }>
    ): Promise<IQuotaApprovalRequest> {
        const { from } = call.request;
        const { info } = call.request;
        const { size } = call.request;
        const { modifiedBy } = call.request;

        return RequestService.create({
            from,
            info,
            size,
            modifiedBy,
        } as IQuotaApprovalRequest);
    }

    static async getRequestByID(
        call: ServerUnaryCall<{ id: string }>
    ): Promise<IQuotaApprovalRequest> {
        const requestId: string = call.request.id;

        return RequestService.getById(requestId);
    }

    static async getRequests(
        call: ServerUnaryCall<GetRequestsQuery>
    ): Promise<IQuotaApprovalRequest[]> {
        const query: GetRequestsQuery = call.request;

        return RequestService.getRequests(query);
    }

    static async updateRequest(
        call: ServerUnaryCall<{
            id: string;
            modifiedBy: string;
            status: string;
        }>
    ): Promise<IQuotaApprovalRequest> {
        const { id } = call.request;
        const { modifiedBy } = call.request;
        const status: RequestStatus = call.request.status as RequestStatus;

        return RequestService.updateRequest(id, modifiedBy, status);
    }
}
