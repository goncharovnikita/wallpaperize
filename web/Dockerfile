FROM node:alpine as deps

COPY package.json /app/package.json

WORKDIR /app

RUN npm install

FROM node:alpine as dist

COPY --from=deps /app/node_modules /app/node_modules

COPY . /app

WORKDIR /app

RUN npm run build:prod

FROM abiosoft/caddy

COPY --from=dist /app/dist /app

COPY Caddyfile /etc/Caddyfile
