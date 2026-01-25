.PHONY: build run clean test

# Build the application
build:
	go build -o resume-api .

# Run the application
run:
	go run .

# Run with hot reload (requires air: go install github.com/cosmtrek/air@latest)
dev:
	air

# Clean build artifacts
clean:
	rm -f resume-api

# Run tests
test:
	go test ./... -v

# Format code
fmt:
	go fmt ./...

# Run linter (requires golangci-lint)
lint:
	golangci-lint run

# Install dependencies
deps:
	go mod download
	go mod tidy
