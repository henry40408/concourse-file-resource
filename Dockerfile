FROM alpine

RUN apk add --no-cache python3

ADD assets/*.py /opt/resource/

RUN mv /opt/resource/check.py /opt/resource/check && \
    mv /opt/resource/in.py /opt/resource/in

RUN chmod +x /opt/resource/check /opt/resource/in
