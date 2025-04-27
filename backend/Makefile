# Makefile for Ecommerce Backend
# This Makefile is used to manage the development environment for the backend

.PHONY: help all build up down migrate test

help:
	@echo "Available commands:"
	@echo "  all     - Builds and starts services"
	@echo "  build   - Build Docker images"
	@echo "  up           - Start all services"
	@echo "  down    - Stop all services"
	@echo "  migrate  - Run migrations (assumes postgres is running)"
	@echo "  test          - Runs the test suite"
	@echo "  clean        - cleans up docker"

all: build up

# Build Docker images
build:
	docker-compose build --no-cache

# Start all services
up:
	docker-compose up -d

# Stop all services
down:
	docker-compose down

migrate:
	docker-compose exec user-service sh -c "goose -dir migrations postgres \"postgres://user:pass@postgres:5432/shopnest?sslmode=disable\" up"

test:
	go test -v ./...

# Clean Docker images
clean:
	docker-compose down --rmi all
	docker system prune -f