FROM node:14.3.0-alpine3.11 as deps

COPY package.json /app/package.json

WORKDIR /app

RUN npm install

FROM node:14.3.0-alpine3.11 as dist

COPY --from=deps /app/node_modules /web/node_modules

COPY . /web

WORKDIR /web

RUN npm run build

FROM abiosoft/caddy

COPY --from=dist /web/build /app

COPY Caddyfile /etc/Caddyfile
