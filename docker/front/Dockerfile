FROM node:14.17.0-alpine

WORKDIR /app

RUN apk update && \
    apk add python make g++ && \
    npm install -g npm && \
    npm install -g vue-cli
