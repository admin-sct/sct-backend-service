.PHONY: build run test clean generate help

# Default target
help:
	@echo "Available targets:"
	@echo "  make build     - Build the server binary"
	@echo "  make run       - Run the server locally"
	@echo "  make generate  - Generate GraphQL code"
	@echo "  make test      - Run tests"
	@echo "  make clean     - Clean build artifacts"

# Build the server
build:
	@go build -o bin/server cmd/server/main.go
	@echo "Build complete: bin/server"

# Run the server locally
run:
	@go run cmd/server/main.go

# Run server with debug mode
debug:
	@go run cmd/server/main.go -debug=true

# Generate GraphQL code
generate:
	@echo "Tidying dependencies..."
	@go mod tidy
	@echo "Generating GraphQL code..."
	@go run github.com/99designs/gqlgen@v0.17.83 generate

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@echo "Clean complete"

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy
	@echo "Dependencies installed"

