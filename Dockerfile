FROM quay.io/henry40408/alpine-node:base-8.9.1

ADD assets/*.js /opt/resource/

RUN mv /opt/resource/check.js /opt/resource/check && \
    mv /opt/resource/in.js /opt/resource/in

RUN chmod +x /opt/resource/check /opt/resource/in
