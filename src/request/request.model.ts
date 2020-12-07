import * as mongoose from "mongoose";
import { IQuotaApprovalRequest } from "./request.interface";

const requestSchema: mongoose.Schema = new mongoose.Schema(
    {
        from: {
            type: String,
            required: true,
        },
        info: {
            type: String,
            required: false,
        },
        status: {
            type: String,
            required: true,
        },
        size: {
            type: Number,
            required: true,
            min: 0,
        },
        modifiedBy: {
            type: String,
            required: false,
        },
    },
    {
        id: true,
        versionKey: false,
        timestamps: true,
        toJSON: {
            virtuals: true,
            transform(doc, ret) {
                delete ret._id;
            },
        },
        toObject: {
            virtuals: true,
            transform(doc, ret) {
                delete ret._id;
            },
        },
    }
);

export const RequestModel = mongoose.model<
    IQuotaApprovalRequest & mongoose.Document
>("QuotaRequest", requestSchema);
