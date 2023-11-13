.PHONY: run
run:
	go run cmd/bot/*

build:
	go build -o bin/bot cmd/bot/*

LOCAL_BIN:=$(CURDIR)/bin
.PHONY: .deps
.deps:
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

MIGRATIONS_DIR=migrations
.PHONY: migration
migration:
	goose -dir=${MIGRATIONS_DIR} create $(NAME) sql

.PHONY: test
test:
	$(info Running tests...)
	go test ./...

.PHONY: cover
cover:
	go test -v $$(go list ./... | grep -v -E './pkg/(api)') -covermode=count -coverprofile=/tmp/c.out
	go tool cover -html=/tmp/c.out

.PHONY: integration
integration:
	$(info Running tests...)
	go test -tags=integration ./tests
