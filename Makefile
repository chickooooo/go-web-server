# Makefile for Go web server

MODULE_NAME := github.com/chickooooo/go-web-server
SERVER_CMD := ./cmd/server

# Show all available commands
.PHONY: help
help: ## Show all available commands
	@grep -E '^[a-zA-Z0-9_-]+:.*?## ' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}' | sort

# Start the web server locally (development mode)
.PHONY: run
run: ## Start the web server locally (development mode)
	@echo "Starting server..."
	go run cmd/server/main.go

# Run all tests in the module
.PHONY: test
test: ## Run all tests in the module
	@echo "Running all tests..."
	go test ./... 

# Run test files in a folder path
# Example: make test-dir dir=internal/api
.PHONY: test-dir
test-dir: ## Run test files in a folder path
ifndef dir
	$(error Please provide a directory using dir=your/dir)
endif
	@echo "Running tests in directory: $(dir)"
	go test ./$(dir)

# Generate coverage report file
.PHONY: cover
cover: ## Generate coverage report file
	@echo "Generating coverage.out ..."
	go test ./... -coverprofile=coverage.out

# View coverage report in terminal (text)
.PHONY: cover-text
cover-text: cover ## View coverage report in terminal (text)
	@echo "Coverage summary:"
	go tool cover -func=coverage.out

# View coverage report in HTML format (opens in browser)
.PHONY: cover-html
cover-html: cover ## View coverage report in HTML format (opens in browser)
	@echo "Opening HTML coverage report..."
	go tool cover -html=coverage.out

# Clean up artifacts
.PHONY: clean
clean: ## Clean up artifacts
	@echo "Cleaning up..."
	rm -f coverage.out
