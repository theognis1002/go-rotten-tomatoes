
### Makefile

```Makefile
# Makefile for the Rotten Tomatoes web crawler project

# The binary name
BINARY_NAME = rotten-tomatoes

# Build the Go binary
build:
	@echo "Building the Go binary..."
	go build -o $(BINARY_NAME) main.go

# Run the Go web crawler
run:
	@echo "Executing the web crawler..."
	./$(BINARY_NAME)


# Clean up build files
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)

# Default target: build the binary
all: build
