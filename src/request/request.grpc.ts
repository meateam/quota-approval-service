import { ServerUnaryCall } from "grpc";
import RequestService from "./request.service";
import {
    IQuotaApprovalRequest,
    GetRequestsQuery,
    RequestStatus,
    QuotaApprovalResponse,
} from "./request.interface";

export default class RequestMethods {
    static async createRequest(
        call: ServerUnaryCall<{
            from: string;
            info: string;
            size: number;
            modifiedBy: string;
        }>
    ): Promise<{ quotaApprovalRequest: any }> {
        const { from } = call.request;
        const { info } = call.request;
        const { size } = call.request;
        const { modifiedBy } = call.request;

        const quotaApprovalRequest: IQuotaApprovalRequest = await RequestService.create(
            {
                from,
                info,
                size,
                modifiedBy,
            } as IQuotaApprovalRequest
        );

        return {
            quotaApprovalRequest: new QuotaApprovalResponse(
                quotaApprovalRequest
            ),
        };
    }

    static async getRequestByID(
        call: ServerUnaryCall<{ id: string }>
    ): Promise<{ quotaApprovalRequest: IQuotaApprovalRequest }> {
        const requestId: string = call.request.id;

        const quotaApprovalRequest = await RequestService.getById(requestId);
        return {
            quotaApprovalRequest: new QuotaApprovalResponse(
                quotaApprovalRequest
            ),
        };
    }

    static async getRequests(
        call: ServerUnaryCall<GetRequestsQuery>
    ): Promise<{ quotaApprovalRequests: IQuotaApprovalRequest[] }> {
        const query: GetRequestsQuery = call.request;

        const quotaApprovalRequests: IQuotaApprovalRequest[] = await RequestService.getRequests(
            query
        );

        return {
            quotaApprovalRequests: quotaApprovalRequests.map(
                (quotaApprovalRequest: IQuotaApprovalRequest) =>
                    new QuotaApprovalResponse(quotaApprovalRequest)
            ),
        };
    }

    static async updateRequest(
        call: ServerUnaryCall<{
            id: string;
            modifiedBy: string;
            status: string;
        }>
    ): Promise<{ quotaApprovalRequest: IQuotaApprovalRequest }> {
        const { id } = call.request;
        const { modifiedBy } = call.request;
        const status: RequestStatus = call.request.status as RequestStatus;

        const quotaApprovalRequest: IQuotaApprovalRequest = await RequestService.updateRequest(
            id,
            modifiedBy,
            status
        );
        return {
            quotaApprovalRequest: new QuotaApprovalResponse(
                quotaApprovalRequest
            ),
        };
    }
}
