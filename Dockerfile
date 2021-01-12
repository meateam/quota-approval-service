FROM node:13.12-alpine
RUN GRPC_HEALTH_PROBE_VERSION=v0.3.0 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe
WORKDIR /usr/src/app

COPY ["package.json", "package-lock.json*", "./"]
RUN npm install --production=true
COPY . .

CMD npm run build && node /usr/src/app/dist/index.js
