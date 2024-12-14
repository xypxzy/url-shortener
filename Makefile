# Variables
BINARY_NAME=url-shortener
CONFIG_PATH=config/local.yaml

# Default target
all: build

# Build the binary
build:
	go build -o $(BINARY_NAME) ./cmd/url-shortener

# Run the application
run:
	CONFIG_PATH=$(CONFIG_PATH) go run ./cmd/url-shortener/main.go

# Clean up generated files
clean:
	rm -f $(BINARY_NAME)

# Format the code
fmt:
	go fmt ./...

# Run tests
test:
	go test ./...

# Display help
help:
	@echo "Usage:"
	@echo "  make build       Build the binary"
	@echo "  make run         Run the application"
	@echo "  make clean       Clean up generated files"
	@echo "  make fmt         Format the code"
	@echo "  make test        Run tests"
	@echo "  make help        Display this help message"
