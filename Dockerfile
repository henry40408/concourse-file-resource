FROM alpine

RUN apk add --no-cache python3

ADD assets/*.py /opt/resource/

RUN chmod +x /opt/resource/*.py
