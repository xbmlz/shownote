FROM golang:alpine

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct
WORKDIR /go/src/shownote
COPY . .

RUN go env && go build -o server .

FROM alpine:latest
LABEL MAINTAINER="997909544@qq.com"

WORKDIR /go/src/shownote

COPY --from=0 /go/src/shownote ./

EXPOSE 8000

ENTRYPOINT ./server -c config.docker.yaml