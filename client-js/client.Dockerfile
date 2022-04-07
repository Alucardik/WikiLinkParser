FROM node:14

WORKDIR /usr/local/client-js
COPY src src
COPY package.json package.json
COPY package-lock.json package-lock.json
COPY src/proto/*.js proto

CMD npm ci
CMD npm run build
