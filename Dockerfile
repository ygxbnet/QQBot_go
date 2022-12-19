FROM golang:1.19-alpine AS builder

RUN go env -w GO111MODULE=auto \
  && go env -w CGO_ENABLED=0 \
  && go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /build

COPY ./ .

RUN set -ex \
    && cd /build \
    && go build -o QQBot_go

FROM alpine:latest

ENV TZ Asia/Shanghai

RUN apk add tzdata && cp /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone \
    && apk del tzdata

COPY --from=builder /build/QQBot_go /usr/bin/QQBot_go
RUN chmod +x /usr/bin/QQBot_go

WORKDIR /data

ENTRYPOINT [ "/usr/bin/QQBot_go" ]
