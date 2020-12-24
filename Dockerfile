FROM golang:1.15.5-alpine3.12

WORKDIR /go

RUN apk add --no-cache git vim
RUN go get -u github.com/labstack/echo