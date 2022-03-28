FROM golang:latest

MAINTAINER YXH

RUN mkdir -p /app/Douban
WORKDIR /app/Douban
COPY . /app/Douban

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn"
RUN go mod download

RUN go build main.go

EXPOSE 8080

ENTRYPOINT  ["./main"]