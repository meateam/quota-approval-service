import { ServerUnaryCall } from "grpc";
import AdminService from "./admin.service";
import { Admin } from "./admin.interface";

export default class AdminMethods {
    static async getAllAdmins(
        _call: ServerUnaryCall<{}>
    ): Promise<{ admins: Admin[] }> {
        const admins = await AdminService.getAllAdmins();
        return { admins };
    }

    static async isUserAdmin(
        call: ServerUnaryCall<{ id: string }>
    ): Promise<{ isUserAdmin: boolean }> {
        const userId: string = call.request.id;

        const isUserAdmin = await AdminService.isUserAdmin(userId);

        return { isUserAdmin };
    }
}
