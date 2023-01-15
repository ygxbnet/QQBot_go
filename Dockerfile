FROM golang:1.19-alpine AS builder

RUN go env -w GO111MODULE=auto \
  && go env -w CGO_ENABLED=0 \
  && go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /QQBot_go

COPY ./ .

RUN set -ex \
    && cd /QQBot_go \
    && go build -o QQBot_go

FROM alpine:latest

RUN apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

COPY --from=builder /build/QQBot_go /usr/bin/QQBot_go
RUN chmod +x /usr/bin/QQBot_go

WORKDIR /data

ENTRYPOINT [ "/usr/bin/QQBot_go" ]
