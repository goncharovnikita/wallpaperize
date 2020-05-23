FROM node:14.3.0-alpine3.11 as deps

COPY package.json /app/package.json

WORKDIR /app

RUN npm install

FROM node:14.3.0-alpine3.11 as dist

COPY --from=deps /app/node_modules /web/node_modules

COPY . /web

WORKDIR /web

RUN npm run build:prod

FROM abiosoft/caddy

COPY --from=dist /web/dist /app

COPY Caddyfile /etc/Caddyfile
