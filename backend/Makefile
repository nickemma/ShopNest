# Makefile for Ecommerce Backend
# This Makefile is used to manage the development environment for the backend

.PHONY: help all build up down migrate test

help:
	@echo "Available commands:"
	@echo "  all     - Builds up"
	@echo "  build   - Build Docker images"
	@echo "  up           - Start all services"
	@echo "  down    - Stop all services"
	@echo "  migrate  - Run migrations (assumes postgres is running)"
	@echo "  test          - Runs the test suite"

all: build up

# Build Docker images
build:
	docker-compose -f docker-compose.yml build --no-cache

# Start all services
up:
	docker-compose -f docker-compose.yml up -d

# Stop all services
down:
	docker-compose -f docker-compose.yml down

migrate:
	docker-compose exec user-service goose -dir /app/migrations postgres "postgres://user:pass@postgres:5432/shopnest?sslmode=disable" up

test:
	go test -v ./...