FROM golang:1.20.3-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY cmd/app/ .

COPY --from=builder /app/ ./
COPY . .
RUN go build -o main ./cmd/app

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/ ./
COPY --from=1 /app/main ./

ENV PORT=8080
EXPOSE $PORT
CMD ["./main"]