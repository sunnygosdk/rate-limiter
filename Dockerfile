FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o rate-limiter ./cmd/server/*

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/rate-limiter .

COPY .env .env

EXPOSE 8080

CMD ["./rate-limiter"]
