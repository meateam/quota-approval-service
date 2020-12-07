import { ServerUnaryCall } from "grpc";
import AdminService from "./admin.service";
import { Admin } from "./admin.interface";

export default class AdminMethods {
    static async getAllAdmins(
        _call: ServerUnaryCall<{}>
    ): Promise<Admin[]> {
        return AdminService.getAllAdmins();
    }

    static async isUserAdmin(
        call: ServerUnaryCall<{ id: string }>
    ): Promise<boolean> {
        const userId: string = call.request.id;

        return AdminService.isUserAdmin(userId);
    }
}
