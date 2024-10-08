GO_BUILD_FLAGS += -ldflags "$(GO_LDFLAGS)"

ifeq ($(GOOS),windows)
	GO_OUT_EXT := .exe
endif

.PHONY: go.fmt
go.fmt: tools.verify.gofumpt tools.verify.goimports
	@$(FIND) -type f -name '*.go' -exec gofumpt -extra -w {} +
	@$(FIND) -type f -name '*.go' -exec goimports -w -local $(ROOT_PACKAGE) {} +
	@go mod edit -fmt

.PHONY: go.lint
go.lint: tools.verify.golangci-lint
	@golangci-lint run -c $(ROOT_DIR)/.golangci.yaml $(ROOT_DIR)/...

.PHONY: go.tidy
go.tidy:
	@go mod tidy

.PHONY: go.test
go.test:
	@mkdir -p $(OUTPUT_DIR)
	@set -o pipefail; \
		go test -race -cover -coverprofile=$(OUTPUT_DIR)/coverage.out -timeout=10m -shuffle=on -short -v $(shell go list ./...) || exit 1
	@if [ "$(shell uname)" = "Darwin" ]; then \
		sed -i '' '/mock_.*.go/d' $(OUTPUT_DIR)/coverage.out; \
		sed -i '' '/internal\/apiserver\/dal\/.*.go/d' $(OUTPUT_DIR)/coverage.out; \
	else \
		sed -i '/mock_.*.go/d' $(OUTPUT_DIR)/coverage.out; \
		sed -i '/internal\/apiserver\/dal\/.*.go/d' $(OUTPUT_DIR)/coverage.out; \
	fi
	@go tool cover -html=$(OUTPUT_DIR)/coverage.out -o $(OUTPUT_DIR)/coverage.html

.PHONY: go.bench
go.bench:
	@go test -bench=".*" -benchmem -benchtime=10s -cpu=4 -timeout=30s

COMMANDS ?= $(filter-out %.md, $(wildcard $(ROOT_DIR)/cmd/*))
BINS ?= $(foreach cmd,${COMMANDS},$(notdir $(cmd)))

ifeq ($(COMMANDS),)
  $(error Could not determine COMMANDS, set ROOT_DIR or run in source dir)
endif
ifeq ($(BINS),)
  $(error Could not determine BINS, set ROOT_DIR or run in source dir)
endif

.PHONY: go.build.verify
go.build.verify:
	@if ! which go &>/dev/null; then echo "Cannot found go compile tool. Please install go tool first."; exit 1; fi

.PHONY: go.build
go.build: go.build.verify $(addprefix go.build., $(addprefix $(PLATFORM)., $(BINS)))

.PHONY: go.build.%
go.build.%:
	$(eval COMMAND := $(word 2,$(subst ., ,$*)))
	$(eval PLATFORM := $(word 1,$(subst ., ,$*)))
	$(eval OS := $(word 1,$(subst _, ,$(PLATFORM))))
	$(eval ARCH := $(word 2,$(subst _, ,$(PLATFORM))))
	@mkdir -p $(OUTPUT_DIR)/platforms/$(OS)/$(ARCH)
	@CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) go build $(GO_BUILD_FLAGS) -o $(OUTPUT_DIR)/platforms/$(OS)/$(ARCH)/$(COMMAND)$(GO_OUT_EXT) $(ROOT_PACKAGE)/cmd/$(COMMAND)

.PHONY: go.air
go.air: tools.verify.air
	@air