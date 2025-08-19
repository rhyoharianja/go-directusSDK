.PHONY: build test clean deps lint fmt vet

# Build the project
build:
	go build -v ./...

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -cover ./...

# Run tests with race detection
test-race:
	go test -v -race ./...

# Install dependencies
deps:
	go mod tidy
	go mod download

# Clean build artifacts
clean:
	go clean -cache -testcache -modcache

# Format code
fmt:
	go fmt ./...

# Run go vet
vet:
	go vet ./...

# Run linter
lint:
	golangci-lint run

# Install development tools
install-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest

# Generate documentation
docs:
	go doc -all ./...

# Run all checks
all: fmt vet lint test

# Build examples
build-examples:
	go build -o bin/basic_example ./examples/basic_usage.go
	go build -o bin/advanced_example ./examples/advanced_usage.go
