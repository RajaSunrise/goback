# test-project-5 Makefile

# Build variables
BINARY_NAME=test-project-5
MAIN_PATH=./cmd/server
BUILD_DIR=./bin
GO_VERSION=$(shell go version | cut -d ' ' -f 3)
GIT_COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
VERSION=0.1.0

# Go variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOMOD=$(GOCMD) mod
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOVET=$(GOCMD) vet
GOFMT=gofmt

# Linker flags
LDFLAGS=-ldflags "-X main.version=$(VERSION) -X main.commit=$(GIT_COMMIT) -X main.buildTime=$(BUILD_TIME)"

# Default target
.PHONY: all
all: clean build

# Help target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build        Build the application"
	@echo "  dev          Run development server with hot reload"
	@echo "  test         Run tests"
	@echo "  test-cover   Run tests with coverage"
	@echo "  test-race    Run tests with race detection"
	@echo "  lint         Run linter"
	@echo "  fmt          Format code"
	@echo "  vet          Run go vet"
	@echo "  clean        Clean build artifacts"
	@echo "  deps         Download dependencies"
	@echo "  deps-update  Update dependencies"
	@echo "  run          Run the application"

# Build the application
.PHONY: build
build: deps
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "Build completed: $(BUILD_DIR)/$(BINARY_NAME)"

# Build for multiple platforms
.PHONY: build-all
build-all: build-linux build-darwin build-windows

.PHONY: build-linux
build-linux:
	@echo "Building for Linux..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)

.PHONY: build-darwin
build-darwin:
	@echo "Building for macOS..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_PATH)

.PHONY: build-windows
build-windows:
	@echo "Building for Windows..."
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)

# Development server with hot reload
.PHONY: dev
dev:
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "Air not found. Install it with: go install github.com/cosmtrek/air@latest"; \
		echo "Running without hot reload..."; \
		$(GOCMD) run $(MAIN_PATH); \
	fi

# Run the application
.PHONY: run
run: build
	@echo "Running $(BINARY_NAME)..."
	$(BUILD_DIR)/$(BINARY_NAME)

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Run tests with coverage
.PHONY: test-cover
test-cover:
	@echo "Running tests with coverage..."
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run tests with race detection
.PHONY: test-race
test-race:
	@echo "Running tests with race detection..."
	$(GOTEST) -race -v ./...

# Run benchmarks
.PHONY: bench
bench:
	@echo "Running benchmarks..."
	$(GOTEST) -bench=. -benchmem ./...

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	$(GOFMT) -l -w .

# Run go vet
.PHONY: vet
vet:
	@echo "Running go vet..."
	$(GOVET) ./...

# Run linter
.PHONY: lint
lint:
	@if command -v golangci-lint > /dev/null; then \
		echo "Running golangci-lint..."; \
		golangci-lint run; \
	else \
		echo "golangci-lint not found. Install it from https://golangci-lint.run/usage/install/"; \
		echo "Running basic checks with go vet..."; \
		$(GOVET) ./...; \
	fi

# Security check
.PHONY: security
security:
	@if command -v gosec > /dev/null; then \
		echo "Running security check..."; \
		gosec ./...; \
	else \
		echo "gosec not found. Install it with: go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest"; \
	fi

# Download dependencies
.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# Update dependencies
.PHONY: deps-update
deps-update:
	@echo "Updating dependencies..."
	$(GOGET) -u ./...
	$(GOMOD) tidy

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html

# Database migrations (GORM)
.PHONY: migrate
migrate:
	@echo "Running database migrations..."
	$(GOCMD) run cmd/migrate/main.go

.PHONY: migrate-down
migrate-down:
	@echo "Rolling back database migrations..."
	$(GOCMD) run cmd/migrate/main.go -down

# Install development tools
.PHONY: install-tools
install-tools:
	@echo "Installing development tools..."
	$(GOINSTALL) github.com/cosmtrek/air@latest
	$(GOINSTALL) github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(GOINSTALL) github.com/securecodewarrior/gosec/v2/cmd/gosec@latest

# Show project info
.PHONY: info
info:
	@echo "Project: $(BINARY_NAME)"
	@echo "Version: $(VERSION)"
	@echo "Go Version: $(GO_VERSION)"
	@echo "Git Commit: $(GIT_COMMIT)"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Framework: Go Gin"
	@echo "Database: PostgreSQL"
	@echo "ORM: GORM"
	@echo "Architecture: Simple Architecture"
	@echo "DevOps: Disabled"

# Create release
.PHONY: release
release: clean test lint build-all
	@echo "Creating release $(VERSION)..."
	@mkdir -p releases
	@tar -czf releases/$(BINARY_NAME)-$(VERSION)-linux-amd64.tar.gz -C $(BUILD_DIR) $(BINARY_NAME)-linux-amd64
	@tar -czf releases/$(BINARY_NAME)-$(VERSION)-darwin-amd64.tar.gz -C $(BUILD_DIR) $(BINARY_NAME)-darwin-amd64
	@zip -j releases/$(BINARY_NAME)-$(VERSION)-windows-amd64.zip $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe
	@echo "Release artifacts created in releases/"

# Development environment setup
.PHONY: setup
setup: deps install-tools
	@echo "Setting up development environment..."
	@if [ ! -f .env ]; then \
		cp .env.example .env; \
		echo "Created .env file from template. Please configure it."; \
	fi
	@echo "Development environment setup complete!"

.DEFAULT_GOAL := help
