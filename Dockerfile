FROM golang:1.15.2-alpine AS build-env
RUN apk update
RUN apk add git

ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8
ENV TZ JST-9
ENV TERM xterm
# Set Environment Variable
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN mkdir /go/src
WORKDIR /go/src
RUN git clone https://github.com/Yuta-K19418/lensent-golang.git
RUN go build -o app main.go