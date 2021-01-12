import mongoose from "mongoose";
import apm from "elastic-apm-node";
import Server, { serviceNames } from "./server";
import config from "./config";
import AdminManager from "./admin/admin.service";
import { HealthCheckResponse } from 'grpc-ts-health-check';

const { mongo, service, admin, apmConfig } = config;

const inititlaizeDB = async () => {
    return AdminManager.createAdmins(admin.list);
};

const initializeMongo = async () => {
    console.log("Connecting to Mongo...");

    await mongoose.connect(mongo.uri, {
        useNewUrlParser: true,
        useUnifiedTopology: true,
        useFindAndModify: false,
        useCreateIndex: true,
    });

    console.log("Mongo connection established");

    await inititlaizeDB();
};

const main = async () => {
    apm.start({
        serviceName: service.name,
        secretToken: apmConfig.secretToken,
        verifyServerCert: apmConfig.verifyServerCert,
        serverUrl: apmConfig.apmURL,
    });

    let isSucceed: boolean = false;

    for (let i = 0; (i < 10) && (!isSucceed); i++) {
        try {
            await initializeMongo();
            isSucceed = true;
        } catch (err) {
            console.error(`Could not connect mongo. Attempt number: ${i + 1}`);
            console.error(err);
            await sleep(1000);
        }
    }

    const quotaApprovalServer: Server = new Server(`${service.port}`);
    
    if (!isSucceed) {
        setHealthStatus(quotaApprovalServer, HealthCheckResponse.ServingStatus.NOT_SERVING);
        console.log(`Server was not created successfully`);
    } else {
        setHealthStatus(quotaApprovalServer, HealthCheckResponse.ServingStatus.SERVING);
        console.log(`Server was created successfully on port ${service.port}`);
    }
    
    quotaApprovalServer.server.start();
};

function setHealthStatus(server: Server, status: number): void {
    for (let i = 0; i < serviceNames.length; i++) {
        server.grpcHealthCheck.setStatus(serviceNames[i], status);
    }
}

function sleep(ms: number) {
    return new Promise((resolve) => setTimeout(resolve, ms));
}

main().catch(console.error);
