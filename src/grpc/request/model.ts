import * as mongoose from 'mongoose';
import { IQuotaApprovalRequest } from './interface';

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
        },
    },
    {
        _id: false,
        id: true,
        timestamps: true,
    },
);

const RequestModel = mongoose.model<IQuotaApprovalRequest & mongoose.Document>('Request', requestSchema);

export default RequestModel;
