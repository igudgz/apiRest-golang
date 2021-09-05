FROM golang:1.6-alpine

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=0N

RUN go get github.com/igudgz/desafioMoneri/src