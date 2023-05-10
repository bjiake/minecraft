FROM golang:1.20.3-alpine

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY cmd/app/ .

RUN go build -o main.app main.go

EXPOSE 8080

ENTRYPOINT ["./main.app"]