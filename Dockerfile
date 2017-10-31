FROM alpine

RUN apk add --no-cache python3

ADD assets/util.py /opt/resource/

ADD assets/check.py /opt/resource/check
ADD assets/in.py /opt/resource/in

RUN chmod +x /opt/resource/*
