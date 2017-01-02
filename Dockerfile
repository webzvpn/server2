FROM kylemanna/openvpn

RUN apk add --update ca-certificates mysql-client make g++ libgcrypt libgcrypt-dev unzip && \
    rm -rf /tmp/* /var/tmp/* /var/cache/apk/*

ADD ./bin /usr/local/bin
RUN chmod a+x /usr/local/bin/*

CMD ["/bin/bash", "-c", "set -e && pre_run"]
