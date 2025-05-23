# Development stage only
FROM golang:1.23-alpine
WORKDIR /app
# Install required system dependencies
RUN apk add --no-cache git curl
# Install air for hot-reloading
RUN go install github.com/air-verse/air@latest && mv /go/bin/air /usr/local/bin/air
# Install goose for migrations
RUN curl -fsSL https://raw.githubusercontent.com/pressly/goose/master/install.sh | sh
ENV PATH="/usr/local/bin:/go/bin:$PATH"
# Copy dependency files first for better caching
COPY go.mod go.sum ./
# Download dependencies with retries
RUN go mod download # Copy source code

COPY . .
EXPOSE 8080
CMD ["air"]

## Build stage for user-service
#FROM golang:1.23-alpine AS builder
#WORKDIR /app
#RUN apk add --no-cache git
#COPY go.mod go.sum ./
#RUN go mod download -x || (echo "go mod download failed with:" && go mod download -x && exit 1)
#COPY . .
#RUN go build -o user-service ./cmd

## Build stage for email-worker
#FROM golang:1.23-alpine AS email-builder
#WORKDIR /app
#RUN apk add --no-cache git
#COPY go.mod go.sum ./
#RUN go mod download -x || (echo "go mod download failed with:" && go mod download -x && exit 1)
#COPY . .
#RUN go build -o email-worker ./email-worker

## Final stage for user-service
#FROM alpine:latest AS user-service
#WORKDIR /app
#RUN apk add --no-cache curl && \
#    echo "Step 1: Installing curl - done" && \
#    echo "Step 2: Downloading goose install script..." && \
#    curl -fsSL https://raw.githubusercontent.com/pressly/goose/master/install.sh -o /tmp/install.sh && \
#    echo "Step 3: Running install script..." && \
#    sh /tmp/install.sh && \
#    echo "Step 4: Verifying goose installation..." && \
#    ls -la /usr/local/bin/goose || (echo "Goose binary not found after install" && exit 1) && \
#    /usr/local/bin/goose --version || (echo "Goose not executable" && exit 1) && \
#    echo "Step 5: Goose installed successfully"
#ENV PATH="/usr/local/bin:$PATH"
#COPY --from=builder /app/user-service .
#COPY ./migrations /app/migrations
#CMD ["sh", "-c", "goose -dir migrations postgres \"$DB_URL\" up && ./user-service"]

## Final stage for email-worker
#FROM alpine:latest AS email-worker
#WORKDIR /app
#COPY --from=email-builder /app/email-worker .
#CMD ["./email-worker"]