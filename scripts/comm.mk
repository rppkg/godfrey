LOG_INFO := echo "[$(shell date +"%Y-%m-%d %H:%M:%S")] INFO:"
LOG_ERROR := echo "[$(shell date +"%Y-%m-%d %H:%M:%S")] ERROR:"

COMMON_MK_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
ROOT_DIR := $(abspath $(shell cd $(COMMON_MK_DIR)/../ && pwd -P))
ROOT_PACKAGE=github.com/rppkg/godfrey
OUTPUT_DIR := $(ROOT_DIR)/_output

ifeq ($(origin TMP_DIR),undefined)
TMP_DIR := $(OUTPUT_DIR)/tmp
$(shell mkdir -p $(TMP_DIR))
endif

FIND := find $(ROOT_DIR) ! -path './third_party/*' ! -path './vendor/*'
XARGS := xargs --no-run-if-empty

## 默认情况下，makefile使用/bin/sh作为shell
SHELL := /bin/bash

PLATFORMS ?= darwin_amd64 windows_amd64 linux_amd64 linux_arm64

ifeq ($(origin PLATFORM), undefined)
	ifeq ($(origin GOOS), undefined)
		GOOS := $(shell go env GOOS)
	endif
	ifeq ($(origin GOARCH), undefined)
		GOARCH := $(shell go env GOARCH)
	endif
	PLATFORM := $(GOOS)_$(GOARCH)
	# 构建镜像时，使用 linux 作为默认的 OS
	IMAGE_PLAT := linux_$(GOARCH)
else
	GOOS := $(word 1, $(subst _, ,$(PLATFORM)))
	GOARCH := $(word 2, $(subst _, ,$(PLATFORM)))
	IMAGE_PLAT := $(PLATFORM)
endif

VERSION_PACKAGE=$(ROOT_DIR)/pkg/version

ifeq ($(origin VERSION), undefined)
VERSION := $(shell git describe --tags --always --match='v*')
endif

GIT_COMMIT:=$(shell git rev-parse HEAD)

GO_VERSION ?= $(shell go version | cut -d ' ' -f 3-)

GO_LDFLAGS += \
	-X $(VERSION_PACKAGE).GitVersion=$(VERSION) \
	-X $(VERSION_PACKAGE).GitCommit=$(GIT_COMMIT) \
	-X $(VERSION_PACKAGE).BuildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')