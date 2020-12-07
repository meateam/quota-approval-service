import { AdminModel } from "./admin.model";
import { Admin } from "./admin.interface";

export default class AdminService {
    static async isUserAdmin(userId: string): Promise<boolean> {
        const admin = await AdminModel.findOne({ id: userId }).exec();

        return !!admin;
    }

    static async getAllAdmins(): Promise<Admin[]> {
        return AdminModel.find().exec();
    }

    static async createAdmins(adminsIds: string[]) {
        return Promise.all(
            adminsIds.map((adminId) => this.createAdmin(adminId))
        );
    }

    static async createAdmin(adminId: string) {
        return AdminModel.findOneAndUpdate(
            { id: adminId },
            { id: adminId },
            { new: true, upsert: true }
        );
    }
}
