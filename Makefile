.DEFAULT_GOAL := all

MAKEFLAGS += --no-print-directory

.PHONY: all
all: gen tidy format lint test build

include scripts/comm.mk
include scripts/tools.mk
include scripts/gen.mk
include scripts/golang.mk

define USAGE_OPTIONS
Options:
  BINS             The binaries to build. Default is all of cmd.
                   This option is available when using: make build
                   Example: make build BINS="godfrey-apiserver godfrey-gincode"
  IMAGES           Backend images to make. Default is all of cmd.
                   This option is available when using: make image/push
                   Example: make image IMAGES="godfrey-apiserver"
  VERSION          The version information compiled into binaries.
                   The default is obtained from gsemver or git.
  V                Set to 1 enable verbose build. Default is 0.
endef
export USAGE_OPTIONS

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

.PHONY: clean
clean:
	@-rm -vrf $(OUTPUT_DIR)