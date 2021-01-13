import { RequestModel } from "./request.model";
import { IQuotaApprovalRequest } from "./request.interface";

export default class RequestRepository {
    static async createRequest(request: IQuotaApprovalRequest) {
        return RequestModel.create(request);
    }

    static async updateRequest(
        requestId: string,
        updateProps: Partial<IQuotaApprovalRequest>
    ) {
        return RequestModel.findOneAndUpdate(
            { _id: requestId },
            {
                $set: updateProps,
            },
            { new: true }
        ).exec();
    }

    static async getRequests(condition: {}) {
        return RequestModel.find(condition).exec();
    }

    static async getRequestById(requestId: string) {
        return RequestModel.findById(requestId).exec();
    }
}
