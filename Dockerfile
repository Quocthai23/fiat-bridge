# Build stage
FROM golang:alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bridge-app ./cmd/bridge/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Install CA certificates for TLS requests
RUN apk --no-cache add ca-certificates tzdata

# Copy the binary from builder
COPY --from=builder /app/bridge-app .

# Expose the API port
EXPOSE 8080

# Command to run the executable
CMD ["./bridge-app"]
