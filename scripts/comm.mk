LOG_INFO := echo "[$(shell date +"%Y-%m-%d %H:%M:%S")] INFO:"
LOG_ERROR := echo "[$(shell date +"%Y-%m-%d %H:%M:%S")] ERROR:"

COMMON_MK_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

ROOT_DIR := $(abspath $(shell cd $(COMMON_MK_DIR)/../ && pwd -P))

ROOT_PACKAGE=github.com/rppkg/godfrey

OUTPUT_DIR := $(ROOT_DIR)/_output

FIND := find . ! -path './third_party/*' ! -path './vendor/*'
XARGS := gxargs --no-run-if-empty