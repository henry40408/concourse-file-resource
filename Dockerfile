FROM alpine

RUN apk add --no-cache jq

ADD assets/* /opt/resource/

RUN chmod +x /opt/resource/*
