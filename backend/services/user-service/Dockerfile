# Build stage for user-service
FROM golang:1.23-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git  # Add git for fetching dependencies
COPY go.mod go.sum ./
RUN go mod download -x || (echo "go mod download failed with:" && go mod download -x && exit 1)
COPY . .
RUN go build -o user-service ./cmd

# Build stage for email-worker
FROM golang:1.23-alpine AS email-builder
WORKDIR /app
RUN apk add --no-cache git  # Add git for fetching dependencies
COPY go.mod go.sum ./
RUN go mod download -x || (echo "go mod download failed with:" && go mod download -x && exit 1)
COPY . .
RUN go build -o email-worker ./email-worker

# Final stage for user-service
FROM alpine:latest AS user-service
WORKDIR /app
COPY --from=builder /app/user-service .
COPY ./migrations /app/migrations
# Install wget and CA certificates
RUN apk add --no-cache wget ca-certificates
# Download goose binary
RUN wget -O /usr/local/bin/goose https://github.com/pressly/goose/releases/download/v3.16.0/goose_linux_x86_64 && \
    chmod +x /usr/local/bin/goose
CMD ["sh", "-c", "goose -dir migrations postgres \"$DB_URL\" up && ./user-service"]

# Final stage for email-worker
FROM alpine:latest AS email-worker
WORKDIR /app
COPY --from=email-builder /app/email-worker .
CMD ["./email-worker"]