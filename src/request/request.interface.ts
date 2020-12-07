export enum RequestStatus {
    PENDING = 'REQUEST_PENDING_APPROVAL',
    APPROVED = 'REQUEST_APPROVED',
    DENIED = 'REQUEST_NOT_APPROVED',
}

export interface IQuotaApprovalRequest {
    id: string;
    status: RequestStatus;
    from: string;
    info: string;
    size: number;
    modifiedBy?: string;
    createdBy: Date;
    updatedBy: Date;
}

export interface GetRequestsQuery {
    createdBy?: string;
    approvableBy?: string;
}
