import mongoose from "mongoose";
import { expect } from "chai";
import { AdminModel } from "./admin.model";
import AdminService from "./admin.service";
import config from "../config";
import { Admin } from "./admin.interface";

describe("Admin Logic", () => {
    before(async () => {
        await mongoose.connection.db.dropDatabase();
        const collections = ["Admins"];
        for (const i in collections) {
            mongoose.connection.db.createCollection(
                collections[i],
                (err) => {}
            );
        }
        // ensure that the model indexes are created
        await AdminModel.ensureIndexes();
    });

    beforeEach(async () => {
        const removeCollectionPromises = [];
        for (const i in mongoose.connection.collections) {
            removeCollectionPromises.push(
                mongoose.connection.collections[i].deleteMany({})
            );
        }
        await Promise.all(removeCollectionPromises);
    });

    describe("#getAdmins", () => {
        it("should return all the admins", async () => {
            const adminId = "-1";
            await AdminService.createAdmin(adminId);
            const admins = await AdminService.getAllAdmins();
            expect(admins).to.exist;
            expect(adminsToIdsArr(admins)).to.eql([adminId]);
        });
    });

    describe("#isUserAdmin", () => {
        it("should return true when user is admin", async () => {
            const adminId = "-1";
            await AdminService.createAdmin(adminId);
            const isAdmin = await AdminService.isUserAdmin(adminId);
            expect(isAdmin).to.exist;
            expect(isAdmin).to.be.true;
        });
    });

    describe("#isUserAdmin", () => {
        it("should return false when user is not admin", async () => {
            const adminId = "-1";
            const notAdminId = "11111";
            await AdminService.createAdmin(adminId);
            const isAdmin = await AdminService.isUserAdmin(notAdminId);
            expect(isAdmin).to.exist;
            expect(isAdmin).to.be.false;
        });
    });

    describe("#createAdmins", () => {
        it("should create all admins", async () => {
            const adminsIds = ["-1", "-2", "-3"];
            const admins = await AdminService.createAdmins(adminsIds);
            expect(admins).to.exist;
            expect(adminsToIdsArr(admins)).to.eql(adminsIds);
        });
    });

    describe("#createAdmin", () => {
        it("should create an admin", async () => {
            const adminId = "-1";
            const admin = await AdminService.createAdmin(adminId);
            expect(admin).to.exist;
            expect(adminToId(admin)).to.equal(adminId);
        });
    });
});

const adminToId = (admin: Admin) => admin.id;
const adminsToIdsArr = (admins: Admin[]) => admins.map(adminToId);
