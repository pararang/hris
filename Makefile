.PHONY: build run test seed

build:
	go build -o bin/api cmd/api/main.go

run:
	go run cmd/main.go

test:
	go test -v ./...

fmt:
	go fmt ./...

tidy:
	go mod tidy

lint:
	golangci-lint run

seed:
	go run cmd/dbseed/main.go
