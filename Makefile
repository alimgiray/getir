.PHONY: build
build:
	go build -trimpath -ldflags="-w -s"  -o ./bin/getir ./cmd/getir/main.go

dev:
	go run ./cmd/getir/main.go

run:
	./bin/getir

all: build run

test: 
	go test ./...