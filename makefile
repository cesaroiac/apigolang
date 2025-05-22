.PHONY: default run run-with-docs build test docs clean

# Variables
APP_NAME=apigolang

# Tasks
default: run-with-docs

run:
	@go run main.go

run-with-docs:
	@swag init
	@go run main.go

build:
	@go build -o $(APP_NAME) main.go

test:
	@go test ./...

docs:
	@swag init

clean:
	@rm -f $(APP_NAME)
	@rm -rf docs
