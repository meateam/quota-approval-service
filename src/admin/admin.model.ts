import * as mongoose from "mongoose";
import { Admin } from "./admin.interface";

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
        versionKey: false,
        toJSON: {
            virtuals: true,
            transform(doc, ret) {
                delete ret._id;
            },
        },
    }
);

export const AdminModel = mongoose.model<Admin & mongoose.Document>(
    "Admin",
    adminSchema
);
