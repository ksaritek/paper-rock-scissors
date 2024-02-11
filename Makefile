LOCAL_BIN := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))/.bin

GOIMPORTS_VERSION=v0.17.0
GRPCURL_VERSION=v1.8.9

export PATH := ${LOCAL_BIN}:$(PATH)

### General
.PHONY: all
all: clean tools build test

.PHONY: tools
tools: go-imports tidy grpcurl-install

### install
.PHONY: go-imports
go-imports:
	GOBIN=$(LOCAL_BIN) go install golang.org/x/tools/cmd/goimports@$(GOIMPORTS_VERSION)

.PHONY: grpcurl-install
grpcurl-install:
	GOBIN=$(LOCAL_BIN) go install github.com/fullstorydev/grpcurl/cmd/grpcurl@$(GRPCURL_VERSION)

### golang utilities
.PHONY: build
build:
	go build -o $(LOCAL_BIN)/server -a -race ./cmd/server
	go build -o $(LOCAL_BIN)/client -a -race ./cmd/client

.PHONY: build-fast
build-fast:
	go build -o $(LOCAL_BIN)/server ./cmd/server
	go build -o $(LOCAL_BIN)/client ./cmd/client


.PHONY: grpc-server-run
grpc-server-run:
	OTEL_EXPORTER_ENABLED=false go run ./cmd/server 

.PHONY: grpc-client-run
grpc-client-run:
	go run ./cmd/client

.PHONY: test
test:
	go test -failfast -v -vet=off -race ./...

.PHONY: update
update: tidy
	go get -u ./...

.PHONY: tidy
tidy:
	go mod tidy -compat=1.21

.PHONY: imports
imports:
	$(LOCAL_BIN)/goimports -l -w .

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: clean
clean:
	rm -rf $(LOCAL_BIN)
