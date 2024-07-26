.DEFAULT_GOAL := all

.PHONY: all
all: gen format lint test build

.PHONY: gen
gen:
	@rm -rf ./internal/apiserver/dal/query
	@go run ./cmd/godfrey-gencode/gencode.go

.PHONY: format
format:

.PHONY: lint
lint:

.PHONY: test
test:

.PHONY: build
build:
