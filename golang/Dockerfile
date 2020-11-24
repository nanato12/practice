FROM golang:1.15-alpine

RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata

ENV GO111MODULE=on

WORKDIR /go/src/workspace
