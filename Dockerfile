# Build stage
FROM golang:1.23-alpine AS builder

# Установка зависимостей для сборки Go-приложения
RUN apk add --no-cache git

# Установка рабочей директории
WORKDIR /app

# Копируем файлы зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Сборка приложения (без CGO, так как Postgres через pgx)
RUN go build -o 2kitchen ./cmd/main.go

# Final stage (чистый рантайм)
FROM alpine:3.18

# Установка минимальных зависимостей
RUN apk add --no-cache ca-certificates

# Копируем бинарник и env-файл
COPY --from=builder /app/2kitchen /2kitchen
COPY --from=builder /app/.env /.env

# Открываем порт (если нужно)
EXPOSE 8080

# Точка входа
CMD ["/2kitchen"]
