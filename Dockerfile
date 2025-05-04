# Build stage
FROM golang:1.23-alpine AS builder

# Install necessary dependencies for building the Go binary and SQLite
RUN apk add --no-cache gcc musl-dev sqlite-dev

# Set the working directory in the container
WORKDIR /app

# Copy go.mod and go.sum for dependency management
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Copy the .env file from the root of the project
COPY .env .env

# Copy the rest of the source code (including main.go from ./cmd)
COPY . .

# Build the Go application with cgo enabled for SQLite support
RUN CGO_ENABLED=1 GOOS=linux go build -o 2kitchen ./cmd/main.go

# Verify that the binary exists in the builder stage
RUN ls -l /app

# Final stage: Use Alpine for the minimal container with dependencies
FROM alpine:3.18

# Install runtime dependencies for SQLite
RUN apk add --no-cache sqlite-libs

# Copy the compiled binary and environment file from the builder stage
COPY --from=builder /app/2kitchen /2kitchen
COPY --from=builder /app/.env /.env

# Verify that the binary exists in the final container image
RUN ls -l /2kitchen

# Set the entry point to the Go binary
CMD ["/2kitchen"]
