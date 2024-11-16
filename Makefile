# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt
GOCOVER=$(GOCMD) tool cover
BINARY_NAME=set-project

# Build flags
LDFLAGS=-ldflags "-w -s"

# Test flags
TEST_FLAGS=-race -v
BENCH_FLAGS=-benchmem
COVER_FLAGS=-coverprofile=coverage.out -covermode=atomic

# Linting
GOLINT=golangci-lint

# Project paths
PKG_PATH=./pkg/...
CMD_PATH=./cmd/...
ALL_PATH=./...

# Colors for terminal output
RED=\033[0;31m
GREEN=\033[0;32m
YELLOW=\033[0;33m
NC=\033[0m # No Color

.PHONY: all build test clean lint fmt help cover bench mod-tidy check install-tools

all: check test build

# Build the project
build:
	@echo "$(GREEN)Building...$(NC)"
	$(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME) $(CMD_PATH)

# Run all tests
test:
	@echo "$(GREEN)Running tests...$(NC)"
	$(GOTEST) $(TEST_FLAGS) $(ALL_PATH)

# Run tests with coverage
cover:
	@echo "$(GREEN)Running tests with coverage...$(NC)"
	$(GOTEST) $(COVER_FLAGS) $(ALL_PATH)
	$(GOCOVER) -func=coverage.out
	$(GOCOVER) -html=coverage.out -o coverage.html
	@echo "$(GREEN)Coverage report generated: coverage.html$(NC)"

# Run benchmarks
bench:
	@echo "$(GREEN)Running benchmarks...$(NC)"
	$(GOTEST) $(BENCH_FLAGS) -run=^$$ -bench=. $(ALL_PATH)

# Clean build artifacts
clean:
	@echo "$(GREEN)Cleaning...$(NC)"
	rm -f coverage.out coverage.html
	rm -f bin/$(BINARY_NAME)
	$(GOCMD) clean -testcache

# Run linter
lint:
	@echo "$(GREEN)Running linter...$(NC)"
	$(GOLINT) run

# Format code
fmt:
	@echo "$(GREEN)Formatting code...$(NC)"
	$(GOFMT) $(ALL_PATH)

# Update dependencies
mod-tidy:
	@echo "$(GREEN)Updating dependencies...$(NC)"
	$(GOMOD) tidy
	$(GOMOD) verify

# Verify project setup
check: mod-tidy lint fmt
	@echo "$(GREEN)All checks passed!$(NC)"


# Show help
help:
	@echo "$(YELLOW)Available targets:$(NC)"
	@echo "  $(GREEN)all$(NC)           - Run checks, tests, and build"
	@echo "  $(GREEN)build$(NC)         - Build the project"
	@echo "  $(GREEN)test$(NC)          - Run tests"
	@echo "  $(GREEN)cover$(NC)         - Run tests with coverage"
	@echo "  $(GREEN)bench$(NC)         - Run benchmarks"
	@echo "  $(GREEN)clean$(NC)         - Clean build artifacts"
	@echo "  $(GREEN)lint$(NC)          - Run linter"
	@echo "  $(GREEN)fmt$(NC)           - Format code"
	@echo "  $(GREEN)mod-tidy$(NC)      - Update dependencies"
	@echo "  $(GREEN)check$(NC)         - Verify project setup"
	@echo "  $(GREEN)help$(NC)          - Show this help message"

# Default target
.DEFAULT_GOAL := help
