.PHONY: go.build.verify
go.build.verify:
	@if ! which go &>/dev/null; then echo "Cannot found go compile tool. Please install go tool first."; exit 1; fi

.PHONY: go.fmt
go.fmt: tools.verify.gofumpt tools.verify.goimports
	@$(FIND) -type f -name '*.go' | $(XARGS) gofumpt -w
	@$(FIND) -type f -name '*.go' | $(XARGS) goimports -w -local $(ROOT_PACKAGE)
	@go mod edit -fmt

.PHONY: go.lint
go.lint: tools.verify.golangci-lint
	@golangci-lint run -c $(ROOT_DIR)/.golangci.yaml $(ROOT_DIR)/...

.PHONY: go.tidy
go.tidy:
	@go mod tidy

.PHONY: go.test
go.test:

.PHONY: go.build
go.build: