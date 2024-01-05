FROM golang:alpine as builder

COPY . .

WORKDIR /go/src/zzk_yunying_task
#作者信息
MAINTAINER "zzk"

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server main.go

#工作目录
WORKDIR /opt
ADD .  /opt

EXPOSE 9000
ENTRYPOINT ./server -c config.test.yaml
