import * as mongoose from 'mongoose';

const adminSchema = new mongoose.Schema(
    {
        id: {
            type: String,
            required: true,
            unique: true,
        },
    },
    {
        _id: false,
    },
);

export const AdminModel = mongoose.model<{ id: string } & mongoose.Document>('Admin', adminSchema);
