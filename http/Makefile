# Set the Go command path and output binary name
GO_CMD=go
BINARY_NAME=myapp
MAIN_PKG=./cmd/myapp/

.PHONY: build run clean

build:
	$(GO_CMD) build -o $(BINARY_NAME) $(MAIN_PKG)

run: build
	set -a && source build/myapp/.env && set +a && ./$(BINARY_NAME)

migrate: build 
	sqlc generate -f ./internal/myapp/db/sqlc.yaml
	set -a && source build/myapp/.env && set +a && ./$(BINARY_NAME) --mode migration
	# goose -dir migrations postgres "user=postgres dbname=myapp sslmode=disable" status

server: build 
	set -a && source build/myapp/.env && set +a && ./$(BINARY_NAME) --mode server

deps:
	docker-compose -f ./build/docker-compose.yaml up --detach

clean:
	rm -f $(BINARY_NAME)
