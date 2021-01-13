import mongoose from "mongoose";
import * as dotenv from "dotenv";
import config from "../config";

const database = config.mongo;

dotenv.config({ path: ".env" });

(<any>mongoose).Promise = Promise;

process.on("unhandledRejection", (reason, p) => {
    console.log("Unhandled Rejection at: Promise", p, "reason:", reason);
    // application specific logging, throwing an error, or other logic here
});

before(async () => {
    const mongoUri = process.env.MONGO_URI || database.uri;
    await mongoose.connect(mongoUri, {
        useCreateIndex: true,
        useNewUrlParser: true,
        useFindAndModify: false,
        useUnifiedTopology: true,
    });
    console.log(`mongo connection: ${mongoUri}`);
});

after((done) => {
    mongoose.connection.close();
    done();
});
