FROM golang:1.20.3-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main.app cmd/app/main.go
ENV docker run --name my-app --network my-network -e MONGO_HOST=mods golang-app-image
EXPOSE 8080

ENTRYPOINT [".cmd/app/main.app"]