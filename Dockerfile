FROM golang:1.20.4-alpine3.18 as builder

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o backend ./cmd/app/main.go

CMD ["./backend"]