TOOLS ?= gofumpt golangci-lint goimports gotests mockgen swagger air

.PHONY: tools.verify
tools.verify: $(addprefix tools.verify., $(TOOLS))

.PHONY: tools.install
tools.install: $(addprefix tools.install., $(TOOLS))

.PHONY: tools.install.%
tools.install.%:
	@$(LOG_INFO) "Starting installation of $*..."
	@$(MAKE) install.$*
	@$(LOG_INFO) "Installation of $* completed successfully."

.PHONY: tools.verify.%
tools.verify.%:
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.$*; fi

.PHONY: install.gofumpt
install.gofumpt:
	@go install mvdan.cc/gofumpt@v0.6.0

.PHONY: install.golangci-lint
install.golangci-lint:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.1
	@golangci-lint completion bash > $(HOME)/.golangci-lint.bash
	@if ! grep -q .golangci-lint.bash $(HOME)/.bashrc; then echo "source \$$HOME/.golangci-lint.bash" >> $(HOME)/.bashrc; fi

.PHONY: install.goimports
install.goimports:
	@go install golang.org/x/tools/cmd/goimports@v0.23.0

.PHONY: install.gotests
install.gotests:
	@go install github.com/cweill/gotests/gotests@v1.6.0

.PHONY: install.mockgen
install.mockgen:
	@go install github.com/golang/mock/mockgen@v1.6.0

.PHONY: install.swagger
install.swagger:
	@go install github.com/go-swagger/go-swagger/cmd/swagger@v0.31.0

.PHONY: install.air
install.air:
	@go install github.com/air-verse/air@v1.52.3
	@air init