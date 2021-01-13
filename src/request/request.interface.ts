export enum RequestStatus {
    PENDING = "REQUEST_PENDING_APPROVAL",
    APPROVED = "REQUEST_APPROVED",
    DENIED = "REQUEST_NOT_APPROVED",
}

export interface IQuotaApprovalRequest {
    id: string;
    status: RequestStatus;
    from: string;
    info?: string;
    size: number;
    modifiedBy?: string;
    createdAt: Date | number;
    updatedAt: Date | number;
}

export interface GetRequestsQuery {
    createdBy?: string;
    approvableBy?: string;
}

export class QuotaApprovalResponse implements IQuotaApprovalRequest {
    id: string;
    status: RequestStatus;
    from: string;
    info?: string | undefined;
    size: number;
    modifiedBy?: string | undefined;
    createdAt: number | Date;
    updatedAt: number | Date;
    
    constructor(quotaApproval: IQuotaApprovalRequest) {
        this.id = quotaApproval.id;
        this.status = quotaApproval.status;
        this.from = quotaApproval.from;
        this.info = quotaApproval.info;
        this.size = quotaApproval.size;
        this.modifiedBy = quotaApproval.modifiedBy;
        this.createdAt = new Date(quotaApproval.createdAt).getTime();
        this.updatedAt = new Date(quotaApproval.updatedAt).getTime();
    }
}
