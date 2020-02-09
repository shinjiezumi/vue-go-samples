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