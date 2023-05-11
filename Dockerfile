FROM golang:1.20.3 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o minecraft ./cmd/app/main.go


CMD [ "./cmd/app/main.go" ]