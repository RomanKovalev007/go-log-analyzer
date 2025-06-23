# Makefile for go-log-analyzer

# Variables
BINARY_NAME := log-analyzer
BUILD_DIR := bin
SOURCE_DIR := ./cmd/analyzer
TEST_DIR := ./...
COVERAGE_FILE := coverage.out

# Build flags
LDFLAGS := -ldflags="-s -w" # Уменьшаем размер бинарника
BUILD_FLAGS := -trimpath

# Default target
all: build

# Build the project
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build $(BUILD_FLAGS) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(SOURCE_DIR)

# Install dependencies
deps:
	@echo "Downloading dependencies..."
	go mod download

# Run tests
test:
	@echo "Running tests..."
	go test -v -cover $(TEST_DIR)

# Run tests with coverage
test-cover:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=$(COVERAGE_FILE) $(TEST_DIR)
	go tool cover -html=$(COVERAGE_FILE) -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run benchmarks
bench:
	@echo "Running benchmarks..."
	go test -bench=. -benchmem $(TEST_DIR)

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -f $(COVERAGE_FILE) coverage.html

# Run the application
run:
	@echo "Running application..."
	go run $(SOURCE_DIR) -f testdata/access.log

# Build for multiple platforms
build-all:
	@echo "Building for multiple platforms..."
	GOOS=linux GOARCH=amd64 make build
	GOOS=windows GOARCH=amd64 make build
	GOOS=darwin GOARCH=amd64 make build

.PHONY: all build deps test test-cover bench lint clean run build-all