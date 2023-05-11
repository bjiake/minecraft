FROM golang:1.20.3 AS builder

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o minecraft ./cmd/app/main.go

CMD["cmd/app/main.go"]