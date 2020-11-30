import { AdminModel } from './model';

export default class AdminService {
    static async isUserAdmin(userId: string) {
        const admin = await AdminModel.findById(userId).exec();

        return !!admin;
    }

    static getAllAdmins() {
        return AdminModel.find().exec();
    }

    static createAdmins(adminsIds: Array<string>) {
        return Promise.all(adminsIds.map((adminId) => this.createAdmin(adminId)));
    }

    static createAdmin(adminId: string) {
        return AdminModel.findOneAndUpdate({ id: adminId }, { id: adminId }, { new: false, upsert: true }).exec();
    }
}
