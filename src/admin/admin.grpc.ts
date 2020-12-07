import { ServerUnaryCall } from 'grpc';
import AdminService from './admin.service';

export default class AdminMethods {
    static async getAllAdmins(_call: ServerUnaryCall<{}>): Promise<Array<{ id: string }>> {
        return AdminService.getAllAdmins();
    }

    static async isUserAdmin(call: ServerUnaryCall<{ id: string }>): Promise<boolean> {
        const userId: string = call.request.id;

        return AdminService.isUserAdmin(userId);
    }
}
