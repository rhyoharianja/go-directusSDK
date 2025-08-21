.PHONY: build test clean lint

# Build the package
build:
	go build ./...

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	go clean

# Run linter
lint:
	go vet ./...

# Install dependencies
deps:
	go mod tidy

# Run example
example:
	go run example_test.go

# Format code
fmt:
	go fmt ./...

# Check for security issues
security:
	go list -json -m all | nancy sleuth

# Generate documentation
docs:
	godoc -http=:6060
