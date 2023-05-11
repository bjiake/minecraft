FROM golang:1.20.3

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -mod=readonly -v -o /go/bin/app ./cmd/app

CMD ["/go/bin/app"]