.DEFAULT_GOAL := goapp

.PHONY: all
all: clean goapp

.PHONY: goapp
goapp:
	mkdir -p bin
	go build -o bin ./...

.PHONY: client
client:
	mkdir -p bin
	go build -o bin ./cmd/client

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	go clean
	rm -f bin/*
	