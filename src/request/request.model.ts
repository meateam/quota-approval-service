import * as mongoose from "mongoose";
import { IQuotaApprovalRequest } from "./request.interface";

const transform = (_doc: any, ret: any) => {
    delete ret._id;
};

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
            transform,
            virtuals: true,
        },
        toObject: {
            transform,
            virtuals: true,
        },
    }
);

export const RequestModel = mongoose.model<
    IQuotaApprovalRequest & mongoose.Document
>("QuotaRequest", requestSchema);
