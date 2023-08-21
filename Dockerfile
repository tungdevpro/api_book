FROM golang:1.20

ENV GO111MODULE=on
WORKDIR /app

COPY . /app

RUN go mod download