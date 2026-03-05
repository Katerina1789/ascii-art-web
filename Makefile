.PHONY: all check build run test coverage clean

# Run formatting and linting
check:
	gofmt -w .
	goimports -w .
	golangci-lint run

# Build and run
build:
	go build -o bin/ascii-art-web ./cmd

run:
	go run ./cmd

# Run all tests with verbose output
test:
	go test -v ./...

# Run tests and generate coverage report
coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func coverage.out

# Clean up generated files
clean:
	rm -rf bin/ coverage.out

# Run the full suite of Pre-PR checks
all: check test coverage
