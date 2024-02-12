FROM golang AS builder

ARG now

RUN go env -w GO111MODULE=auto \
  && go env -w CGO_ENABLED=0 \
  && go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /QQBot_go

COPY ./ .

RUN set -ex \
  && cd /QQBot_go \
  && go build -ldflags "-X 'QQBot_go/internal/base.BuildTime=${now}'" -o QQBot_go


FROM alpine:latest

COPY --from=builder /QQBot_go/QQBot_go /usr/bin/QQBot_go

RUN apk update \
  && apk add tzdata --no-cache \
  && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
  && echo "Asia/Shanghai" > /etc/timezone \
  && chmod +x /usr/bin/QQBot_go

WORKDIR /data

ENTRYPOINT [ "/usr/bin/QQBot_go" ]
