.PHONY: help test lint fmt vet security check

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "}; /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

test: ## Run all tests
	go test -race -coverprofile=coverage.txt -covermode=atomic -v ./...

fmt: ## Run code formatting
	go fmt ./...

vet: ## Run static code analysis
	go vet ./...

lint: vet fmt ## Run linting/formatting tools
	golangci-lint run

security: ## Run security scans
	gosec ./...
	govulncheck ./...

check: fmt lint test security ## Run all checks: formatting, linting, tests, and security
