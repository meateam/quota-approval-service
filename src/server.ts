import * as grpc from 'grpc';
import RequestMethods from './request/request.grpc';
import AdminMethods from './admin/admin.grpc';
import { loadSync } from '@grpc/proto-loader';
import { wrapper } from './utils/logger';

export const serviceNames: string[] = ['', 'users.Users'];

const QUOTA_APPROVAL_PROTO_PATH = `${__dirname}/../proto/quotaApproval/quotaApproval.proto`;
const ADMIN_PROTO_PATH = `${__dirname}/../proto/admin/admin.proto`;

const quotaApprovalPackageDefinition = loadSync(QUOTA_APPROVAL_PROTO_PATH, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
});

const adminPackageDefinition = loadSync(ADMIN_PROTO_PATH, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
});

const quotaApprovalProto: any = grpc.loadPackageDefinition(quotaApprovalPackageDefinition).quotaApproval;
const adminProto: any = grpc.loadPackageDefinition(adminPackageDefinition).admins;

export default class RPC {
    public server: any;

    public constructor(port: string) {
        this.server = new grpc.Server();

        this.server.addService(quotaApprovalProto.QuotaApproval.service, {
            CreateQuotaApproval: wrapper(RequestMethods.createRequest, 'CreateQuotaApproval'),
            GetQuotasApprovals: wrapper(RequestMethods.getRequests, 'GetQuotasApprovals'),
            GetQuotaApprovalById: wrapper(RequestMethods.getRequestById, 'GetQuotaApprovalById'),
            UpdateQuotaApproval: wrapper(RequestMethods.updateRequest, 'UpdateQuotaApproval'),
        });

        this.server.addService(adminProto.Admins.service, {
            GetAllAdmins: wrapper(AdminMethods.getAllAdmins, 'GetAllAdmins'),
            IsUserAdmin: wrapper(AdminMethods.isUserAdmin, 'IsUserAdmin'),
        });

        this.server.bind(`0.0.0.0:${port}`, grpc.ServerCredentials.createInsecure());
    }
}
