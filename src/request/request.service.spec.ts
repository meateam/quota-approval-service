import mongoose from "mongoose";
import AdminManager from "../admin/admin.service";
import RequestService from "./request.service";
import { expect } from "chai";
import { RequestModel } from "./request.model";
import { IQuotaApprovalRequest, RequestStatus } from "./request.interface";
import config from "../config";
const { createMockServer } = require("grpc-mock");

const mockServer = createMockServer({
    protoPath: `${__dirname}/../../proto/quota/quota.proto`,
    packageName: "quota",
    serviceName: "QuotaService",
    rules: [
        {
            method: "ChangeQuotaLimit",
            input: ".*",
            output: { message: "Quota updated successfully!" },
        },
    ],
});

const port = config.quotaService.url.split(":")[1];

describe("Quota approval Logic", () => {
    before(async () => {
        await mongoose.connection.db.dropDatabase();
        const collections = ["QuotaRequest", "Admins"];
        for (const i in collections) {
            mongoose.connection.db.createCollection(
                collections[i],
                (err) => {}
            );
        }
        // ensure that the model indexes are created
        await RequestModel.ensureIndexes();

        mockServer.listen(`0.0.0.0:${port}`);
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

    after(() => {
        mockServer.close(true);
    });

    afterEach(() => mockServer.clearInteractions());

    describe("#getById", () => {
        it("should return a quota approval request", async () => {
            let quotaRequest: IQuotaApprovalRequest = {
                size: 2,
                from: "-1",
                info: "some info",
            } as IQuotaApprovalRequest;
            quotaRequest = await RequestService.create(quotaRequest);
            const result = await RequestService.getById(quotaRequest.id);
            expect(result).to.exist;
            expect(result.status).to.be.equal(RequestStatus.PENDING);
            expect(result.from).to.be.equal("-1");
            expect(result.info).to.be.equal("some info");
            expect(result.size).to.be.equal(2);
            expect(result.modifiedBy).to.not.exist;
            expect(result.createdAt instanceof Date).to.be.true;
            expect(result.updatedAt instanceof Date).to.be.true;
        });
    });

    describe("#create", () => {
        it("should create a quota approval request", async () => {
            const quotaRequest: IQuotaApprovalRequest = {
                size: 2,
                from: "-1",
                info: "some info",
            } as IQuotaApprovalRequest;
            const result = await RequestService.create(quotaRequest);
            expect(result).to.exist;
            expect(result.status).to.be.equal(RequestStatus.PENDING);
            expect(result.from).to.be.equal("-1");
            expect(result.info).to.be.equal("some info");
            expect(result.size).to.be.equal(2);
            expect(result.modifiedBy).to.not.exist;
            expect(result.createdAt instanceof Date).to.be.true;
            expect(result.updatedAt instanceof Date).to.be.true;
        });
    });

    describe("#getRequests", () => {
        it("should return all quota approval requests", async () => {
            const quotaRequest: IQuotaApprovalRequest = {
                size: 2,
                from: "-1",
                info: "some info",
            } as IQuotaApprovalRequest;
            await Promise.all([
                RequestService.create(quotaRequest),
                RequestService.create(quotaRequest),
            ]);
            const result = await RequestService.getRequests({});
            expect(result).to.exist;
            expect(result.length).to.equal(2);
            result.forEach((quotaReq: IQuotaApprovalRequest) => {
                expect(quotaReq.status).to.be.equal(RequestStatus.PENDING);
                expect(quotaReq.from).to.be.equal("-1");
                expect(quotaReq.info).to.be.equal("some info");
                expect(quotaReq.size).to.be.equal(2);
                expect(quotaReq.modifiedBy).to.not.exist;
                expect(quotaReq.createdAt instanceof Date).to.be.true;
                expect(quotaReq.updatedAt instanceof Date).to.be.true;
            });
        });
    });

    describe("#getRequests", () => {
        it("should return all quota approval requests that can be approved by admin", async () => {
            await AdminManager.createAdmin("-2");
            const quotaRequest: IQuotaApprovalRequest = {
                size: 2,
                from: "-1",
                info: "some info",
            } as IQuotaApprovalRequest;
            await Promise.all([
                RequestService.create(quotaRequest),
                RequestService.create(quotaRequest),
            ]);
            const result = await RequestService.getRequests({
                approvableBy: "-2",
            });
            expect(result).to.exist;
            expect(result.length).to.equal(2);
            result.forEach((quotaReq: IQuotaApprovalRequest) => {
                expect(quotaReq.status).to.be.equal(RequestStatus.PENDING);
                expect(quotaReq.from).to.be.equal("-1");
                expect(quotaReq.info).to.be.equal("some info");
                expect(quotaReq.size).to.be.equal(2);
                expect(quotaReq.modifiedBy).to.not.exist;
                expect(quotaReq.createdAt instanceof Date).to.be.true;
                expect(quotaReq.updatedAt instanceof Date).to.be.true;
            });
        });
    });

    describe("#getRequests", () => {
        it("should return all quota approval requests that can be approved by non admin", async () => {
            const quotaRequest: IQuotaApprovalRequest = {
                size: 2,
                from: "-1",
                info: "some info",
            } as IQuotaApprovalRequest;
            await Promise.all([
                RequestService.create(quotaRequest),
                RequestService.create(quotaRequest),
            ]);
            const result = await RequestService.getRequests({
                approvableBy: "-1",
            });
            expect(result).to.exist;
            expect(result.length).to.equal(0);
            expect(result).to.be.eql([]);
        });
    });

    describe("#getRequests", () => {
        it("should return all quota approval requests that can be approved by admin", async () => {
            const quotaRequest: IQuotaApprovalRequest = {
                size: 2,
                from: "-1",
                info: "some info",
            } as IQuotaApprovalRequest;
            await Promise.all([
                RequestService.create(quotaRequest),
                RequestService.create(quotaRequest),
            ]);
            const result = await RequestService.getRequests({
                createdBy: "-1",
            });
            expect(result).to.exist;
            expect(result.length).to.equal(2);
            result.forEach((quotaReq: IQuotaApprovalRequest) => {
                expect(quotaReq.status).to.be.equal(RequestStatus.PENDING);
                expect(quotaReq.from).to.be.equal("-1");
                expect(quotaReq.info).to.be.equal("some info");
                expect(quotaReq.size).to.be.equal(2);
                expect(quotaReq.modifiedBy).to.not.exist;
                expect(quotaReq.createdAt instanceof Date).to.be.true;
                expect(quotaReq.updatedAt instanceof Date).to.be.true;
            });
        });
    });

    describe("#updateRequest", () => {
        it("should approve quota approval request", async () => {
            await AdminManager.createAdmin("-2");
            let quotaRequest: IQuotaApprovalRequest = {
                size: 2,
                from: "-1",
                info: "some info",
            } as IQuotaApprovalRequest;
            quotaRequest = await RequestService.create(quotaRequest);
            const result = await RequestService.updateRequest(
                quotaRequest.id,
                "-2",
                RequestStatus.APPROVED
            );
            expect(result).to.exist;
            expect(result.status).to.be.equal(RequestStatus.APPROVED);
            expect(result.from).to.be.equal("-1");
            expect(result.info).to.be.equal("some info");
            expect(result.size).to.be.equal(2);
            expect(result.modifiedBy).to.be.equal("-2");
            expect(result.createdAt instanceof Date).to.be.true;
            expect(result.updatedAt instanceof Date).to.be.true;
        });
    });

    describe("#updateRequest", () => {
        it("should throw UnauthorizedError for non admin", async () => {
            let quotaRequest: IQuotaApprovalRequest = {
                size: 2,
                from: "-1",
                info: "some info",
            } as IQuotaApprovalRequest;
            quotaRequest = await RequestService.create(quotaRequest);
            let updateRequestError: Error = new Error();
            try {
                const result = await RequestService.updateRequest(
                    quotaRequest.id,
                    "-1",
                    RequestStatus.APPROVED
                );
            } catch (err) {
                updateRequestError = err;
            }
            expect(updateRequestError).to.be.an("Error");
            expect(updateRequestError.message).to.equal(
                `Operation not permitted for user: -1`
            );
        });
    });
});
