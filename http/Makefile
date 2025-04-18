# Set the Go command path and output binary name
GO_CMD=go
BINARY_NAME=myapp
MAIN_PKG=./cmd/myapp/

.PHONY: build run clean

build:
	$(GO_CMD) build -o $(BINARY_NAME) $(MAIN_PKG)

run: build
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)
