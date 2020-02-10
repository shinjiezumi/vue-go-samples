FROM golang:1.13.7

RUN apt-get update && apt-get install -y --no-install-recommends \
    git \
    zip \
    unzip \
    iputils-ping \
    vim

COPY ./src/api /go/src/github.com/shinjiezumi/vue-go-samples/src/api

WORKDIR /go/src/github.com/shinjiezumi/vue-go-samples/src/api

RUN go get -u github.com/codegangsta/gin \
  && go get -u github.com/golang/dep/cmd/dep

ENV DBMS="mysql" \
    USER="root" \
    PASS="root" \
    PROTOCOL="tcp(db:3306)" \
    DBNAME="sample" \
    JWT_SECRET_KEY="75c92a074c341e9964329c0550c2673730ed8479c885c43122c90a2843177d5ef21cb50cfadcccb20aeb730487c11e09ee4dbbb02387242ef264e74cbee97213"