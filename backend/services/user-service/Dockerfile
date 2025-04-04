# Build stage for user-service
FROM golang:1.23-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download -x || (echo "go mod download failed with:" && go mod download -x && exit 1)
COPY . .
RUN go build -o user-service ./cmd

# Build stage for email-worker
FROM golang:1.23-alpine AS email-builder
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download -x || (echo "go mod download failed with:" && go mod download -x && exit 1)
COPY . .
RUN go build -o email-worker ./email-worker

# Final stage for user-service
FROM alpine:latest AS user-service
WORKDIR /app
RUN apk add --no-cache curl && \
    echo "Step 1: Installing curl - done" && \
    echo "Step 2: Downloading goose install script..." && \
    curl -fsSL https://raw.githubusercontent.com/pressly/goose/master/install.sh -o /tmp/install.sh && \
    echo "Step 3: Running install script..." && \
    sh /tmp/install.sh && \
    echo "Step 4: Verifying goose installation..." && \
    ls -la /usr/local/bin/goose || (echo "Goose binary not found after install" && exit 1) && \
    /usr/local/bin/goose --version || (echo "Goose not executable" && exit 1) && \
    echo "Step 5: Goose installed successfully"
ENV PATH="/usr/local/bin:$PATH"
COPY --from=builder /app/user-service .
COPY ./migrations /app/migrations
CMD ["sh", "-c", "goose -dir migrations postgres \"$DB_URL\" up && ./user-service"]

# Final stage for email-worker
FROM alpine:latest AS email-worker
WORKDIR /app
COPY --from=email-builder /app/email-worker .
CMD ["./email-worker"]