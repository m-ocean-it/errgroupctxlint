.PHONY: build
build:
	mkdir -p ./bin
	go build -o ./bin/errgroupctxlint ./cmd/errgroupctxlint

.PHONY: run
run: build
	./bin/errgroupctxlint

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: check
check: build test lint

.PHONY: clean
clean:
	rm -rf ./bin
