.PHONY: help backend frontend install clean dev build docker-up docker-down docker-logs docker-rebuild docker-clean

help:
	@echo "Available commands:"
	@echo ""
	@echo "Local Development:"
	@echo "  make install        - Install all dependencies (backend and frontend)"
	@echo "  make dev            - Run both backend and frontend in development mode"
	@echo "  make backend        - Run only the backend API"
	@echo "  make frontend       - Run only the frontend"
	@echo "  make build          - Build both backend and frontend"
	@echo "  make clean          - Clean build artifacts"
	@echo ""
	@echo "Docker Commands:"
	@echo "  make docker-up      - Start all services with Docker Compose"
	@echo "  make docker-down    - Stop all Docker services"
	@echo "  make docker-logs    - View logs from all services"
	@echo "  make docker-rebuild - Rebuild and restart all services"
	@echo "  make docker-clean   - Remove all containers, images, and volumes"

install:
	@echo "Installing backend dependencies..."
	cd backend && go mod download && go mod tidy
	@echo "Installing frontend dependencies..."
	cd frontend && npm install

dev:
	@echo "Starting backend and frontend..."
	@make -j2 backend frontend

backend:
	@echo "Starting backend API on http://localhost:8080"
	cd backend && go run .

frontend:
	@echo "Starting frontend on http://localhost:3000"
	cd frontend && npm run dev

build:
	@echo "Building backend..."
	cd backend && go build -o resume-api .
	@echo "Building frontend..."
	cd frontend && npm run build

clean:
	@echo "Cleaning build artifacts..."
	cd backend && rm -f resume-api
	cd frontend && rm -rf .next out node_modules/.cache

# Docker Compose Commands
docker-up:
	@echo "Starting all services with Docker Compose..."
	docker compose up -d
	@echo "Services started. Frontend available at http://localhost:3000"

docker-down:
	@echo "Stopping all Docker services..."
	docker compose down

docker-logs:
	@echo "Viewing logs (Ctrl+C to exit)..."
	docker compose logs -f

docker-rebuild:
	@echo "Rebuilding and restarting all services..."
	docker compose up -d --build

docker-clean:
	@echo "Cleaning up all Docker resources..."
	docker compose down --rmi all --volumes
	@echo "Docker cleanup complete"
