FROM golang:alpine as builder

COPY . .

WORKDIR /go/src/zzk_yunying_task
#作者信息
MAINTAINER "zzk"

RUN go generate && go env && go build -o server .

#工作目录
WORKDIR /opt
ADD .  /opt

EXPOSE 9000
ENTRYPOINT ./server -c config.test.yaml
