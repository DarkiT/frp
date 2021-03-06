FROM golang:1.8

COPY . /go/src/github.com/fatedier/frp

RUN cd /go/src/github.com/fatedier/frp \
 && make \
 && mv bin/frpc /frpc \
 && mv bin/frps /frps \
 && mv conf/frpc.ini /frpc.ini \
 && mv conf/frps.ini /frps.ini \
 && mv conf/init.sh /init.sh \
 && make clean

RUN chmod +x /init.sh
VOLUME /data
WORKDIR /

EXPOSE 80 443 6000 7000 7500

ENTRYPOINT ["/init.sh"]
