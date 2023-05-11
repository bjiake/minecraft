FROM golang:1.20.3-alpine

ENV HOST=0.0.0.0

WORKDIR /app
#RUN mkdir /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main cmd/app/main.go

EXPOSE 8080

ENTRYPOINT ["/app/main"]