import * as env from 'env-var';
import * as dotenv from 'dotenv';

const dotenvPath = process.env.LOAD_DEV_DOTENV ? '.env.dev' : '.env';

dotenv.config({ path: dotenvPath });

const esHost = env.get('ELASTICSEARCH_URL').default('http://localhost:9200').asArray(',');
const esUser = env.get('ELASTICSEARCH_USER').default('');
const esPass = env.get('ELASTICSEARCH_PASSWORD').default('');

const config = {
    service: {
        port: env.get('PORT').required().asPortNumber(),
        name: env.get('QAS_APM_SERVICE_NAME').default('quota-approval-service').asString(),
    },
    mongo: {
        uri: env.get('MONGO_URI').default('mongodb://localhost:27017/kdrive').required().asUrlString(),
    },
    quotaService: {
        url: env.get('QUOTA_SERVICE_URL').required().asUrlString(),
    },
    admin: {
        list: env.get('ADMINS_ID_LIST').required().asArray(','),
    },
    apmConfig: {
        secretToken: env.get('APM_SECRET_TOKEN').default('').asString(),
        verifyServerCert: env.get('ELASTIC_APM_VERIFY_SERVER_CERT').default('false').asBool(),
        apmURL: env.get('ELASTIC_APM_SERVER_URL').default('http://localhost:8200').asUrlString(),
    },
    elasticsearch: {
        esHost,
        esUser,
        esPass,
    },
    debugMode: env.get('DEBUG_MODE').default('false').asBool(),
    confLogger: {
        options: {
            hosts: esHost,
            // Might be auth instead, not sure.
            httpAuth: `${esUser}:${esPass}`,
        },
        indexPrefix: env.get('LOG_INDEX').default('kdrive').asString(),
    },
};

export default config;
