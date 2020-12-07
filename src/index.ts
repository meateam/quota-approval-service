/* eslint-disable no-console */
import * as mongoose from 'mongoose';
import Server from './server';
import config from './config';
import AdminManager from './admin/admin.service';
import * as apm from 'elastic-apm-node';

const { mongo, service, admin, apmConfig } = config;

const inititlaizeDB = async () => {
    return AdminManager.createAdmins(admin.list);
};

const initializeMongo = async () => {
    console.log('Connecting to Mongo...');

    await mongoose.connect(mongo.uri, { useNewUrlParser: true, useUnifiedTopology: true, useFindAndModify: false, useCreateIndex: true });

    console.log('Mongo connection established');

    await inititlaizeDB();
};

const main = async () => {
    apm.start({
        serviceName: service.name,
        secretToken: apmConfig.secretToken,
        verifyServerCert: apmConfig.verifyServerCert,
        serverUrl: apmConfig.apmURL,
    });

    await initializeMongo();

    new Server("" + service.port);
};

main().catch(console.error);
