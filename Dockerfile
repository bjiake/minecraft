FROM golang:1.20.3-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main cmd/app/main.go

EXPOSE 8080

ENTRYPOINT ["/app/cmd/app/main"]