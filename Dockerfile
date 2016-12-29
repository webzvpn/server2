FROM kylemanna/openvpn

RUN apk add --update mysql-client make g++ libgcrypt libgcrypt-dev unzip && 
    rm -rf /tmp/* /var/tmp/* /var/cache/apk/*

ADD ./gcsdownloader /usr/local/bin/gcsdownloader
RUN chmod a+x /usr/local/bin/gcsdownloader

ADD ./pre_run /usr/local/bin/pre_run
RUN chmod a+x /usr/local/bin/pre_run

CMD ["pre_run"]
