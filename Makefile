# Binary name
BINARY=product-service

# Go command
GO=go

# Run the application
run:
	$(GO) run cmd/main.go

# Run tests
test:
	$(GO) test ./... -v

# Format the code
fmt:
	$(GO) fmt ./...

# Run linting (requires golangci-lint)
lint:
	golangci-lint run ./...

# Clean up any build artifacts
clean:
	rm -f $(BINARY)

# Build the binary
build:
	$(GO) build -o $(BINARY) cmd/main.go

# Run everything (format, lint, test, build)
all: fmt lint test build

