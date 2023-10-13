.PHONY: all

BINARY=simple-worker

all: test build lint

test:
	go test -race -v ./...
build:
	go build -ldflags "-s -w" -o $(BINARY) -v ./cmd/$(BINARY)

build-debug:
	go build -race -gcflags "all=-N -l" -o $(BINARY) -v ./cmd/$(BINARY)

clean:
	go clean
	rm -f $(BINARY)
fmt:
	goimports -w .
	gofmt -w .

lint:
	golangci-lint run