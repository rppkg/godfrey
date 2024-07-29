.DEFAULT_GOAL := all

MAKEFLAGS += --no-print-directory

.PHONY: all
all: gen tidy format lint test build

include scripts/comm.mk
include scripts/tools.mk
include scripts/gen.mk
include scripts/golang.mk

.PHONY: gen
gen:
	@$(MAKE) gen.query

.PHONY: tidy
tidy:
	@$(MAKE) go.tidy

.PHONY: format
format:
	@$(MAKE) go.fmt

.PHONY: lint
lint:
	@$(MAKE) go.lint

.PHONY: test
test:
	@$(MAKE) go.test

.PHONY: build
build:
	@$(MAKE) go.build
