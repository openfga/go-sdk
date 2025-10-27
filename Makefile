.PHONY: help test test-unit test-integration lint fmt vet security

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "}; /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

test: test-unit ## Run all tests

test-unit: ## Run unit tests only
	go test -race -coverprofile=coverage.txt -covermode=atomic -v ./...

test-integration: ## Run integration tests only (if available, tagged with *integration*)
	go test -v -tags=integration -count=1 ./...

fmt: ## Run code formatting
	go fmt ./...

vet: ## Run static code analysis
	go vet ./...

lint: vet fmt ## Run linting/formatting tools
	golangci-lint run

security: ## Run security scans
	gosec ./...
	govulncheck ./...

check: fmt lint test security # Run all and everything
