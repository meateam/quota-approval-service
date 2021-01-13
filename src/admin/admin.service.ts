import AdminRepository from "./admin.repository";
import { Admin } from "./admin.interface";

export default class AdminService {
    static async isUserAdmin(userId: string): Promise<boolean> {
        const admin = await AdminRepository.getAdminById(userId);

        return !!admin;
    }

    static async getAllAdmins(): Promise<Admin[]> {
        return AdminRepository.getAdmins({});
    }

    static async createAdmins(adminsIds: string[]) {
        return Promise.all(
            adminsIds.map((adminId) => this.createAdmin(adminId))
        );
    }

    static async createAdmin(adminId: string) {
        return AdminRepository.createAdmin(adminId);
    }
}
