FROM golang:alpine

WORKDIR /golang-docker

ADD . .
COPY go.mod go.sum ./
RUN go mod download

ENTRYPOINT go build  && ./golang-docker