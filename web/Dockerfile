FROM node:alpine as deps

COPY package.json /app/package.json

WORKDIR /app

RUN npm install

FROM node:alpine as dist

COPY --from=deps /app/node_modules /web/node_modules

COPY . /web

WORKDIR /web

RUN npm run build:prod

FROM abiosoft/caddy

COPY --from=dist /web/dist /app

COPY Caddyfile /etc/Caddyfile
