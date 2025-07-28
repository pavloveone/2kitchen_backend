FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o 2kitchen ./cmd/main.go

FROM alpine:3.18

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/2kitchen /2kitchen
COPY --from=builder /app/.env /.env

EXPOSE 8080

CMD ["/2kitchen"]
