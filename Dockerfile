FROM golang:1.8

VOLUME /data

COPY . /go/src/github.com/fatedier/frp

RUN cd /go/src/github.com/fatedier/frp \
 && make \
 && cp bin/frpc /data/frpc \
 && cp bin/frps /data/frps \
 && mv bin/frpc /frpc \
 && mv bin/frps /frps \
 && mv conf/frpc.ini /frpc.ini \
 && mv conf/frps.ini /frps.ini \
 && make clean

WORKDIR /

EXPOSE 80 443 6000 7000 7500

ENTRYPOINT ["/frps"]
