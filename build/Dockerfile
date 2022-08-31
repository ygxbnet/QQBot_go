FROM golang:1.17-alpine As bnilder

RUN go env -w GO111MODULE=auto \
    && go env -w CGO_ENABLED=0 \
    && go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /build

COPY .. .

RUN cd /build \
    && go build -o qqbot-go

FROM alpine

COPY --from=bnilder /build/qqbot-go /usr/bin/
# RUN chmod +x /usr/bin/qqbot-go

WORKDIR /data

CMD ["/usr/bin/qqbot-go"]