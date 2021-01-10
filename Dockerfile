FROM node:13.12-alpine

WORKDIR /usr/src/app

COPY ["package.json", "package-lock.json*", "./"]
RUN npm install --production=true
COPY . .

# RUN chmod g+rwx -R /usr/src/app

CMD npm run build && node /usr/src/app/dist/index.js
# CMD npm run dev