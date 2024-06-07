# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=email-validator-service
BINARY_UNIX=$(BINARY_NAME)_unix

# Directories
CMD_DIR=cmd/server
TEST_DIR=./...

# Default target executed when no arguments are given to make
all: test build

# Build the binary
build:
	$(GOBUILD) -o $(BINARY_NAME) $(CMD_DIR)

# Run the binary
run:
	$(GORUN) $(CMD_DIR)/main.go

# Run tests
test:
	$(GOTEST) -v $(TEST_DIR)

# Clean the binary and built files
clean:
	if [ -f $(BINARY_NAME) ]; then rm $(BINARY_NAME); fi
	if [ -f $(BINARY_UNIX) ]; then rm $(BINARY_UNIX); fi

# Build the binary for Linux
build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) $(CMD_DIR)

# Install dependencies
deps:
	$(GOGET) -u github.com/gorilla/mux
	$(GOGET) -u github.com/sufimalek/emailvalidator

.PHONY: all build clean run test build-linux deps
