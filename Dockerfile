FROM golang:1.9-alpine AS builder

COPY . /go/src/github.com/henry40408/concourse-file-resource

RUN apk --no-cache add make && \
    cd /go/src/github.com/henry40408/concourse-file-resource && \
    make build-linux

WORKDIR /opt/resource

FROM alpine:edge

COPY --from=builder /opt/resource /opt/resource
