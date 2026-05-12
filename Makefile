.PHONY: build run test lint check clean

SRC := $(shell find . -type f -name '*.go' ! -name '*_test.go')

build: ./bin/errgroupctxlint

run: build
	./bin/errgroupctxlint

test:
	go test ./...

lint:
	golangci-lint run

check: build test lint

clean:
	rm -rf ./bin

./bin/errgroupctxlint: $(SRC) | ./bin
	go build -o ./bin/errgroupctxlint ./cmd/errgroupctxlint

./bin:
	mkdir -p ./bin
