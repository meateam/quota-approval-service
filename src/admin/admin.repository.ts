import { AdminModel } from "./admin.model";

export default class AdminRepository {
    static async createAdmin(adminId: string) {
        return AdminModel.findOneAndUpdate(
            { id: adminId },
            { id: adminId },
            { new: true, upsert: true }
        );
    }

    static async getAdmins(condition: {}) {
        return AdminModel.find(condition).exec();
    }

    static async getAdminById(adminId: string) {
        return AdminModel.findOne({ id: adminId }).exec();
    }
}
