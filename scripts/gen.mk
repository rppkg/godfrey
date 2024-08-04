.PHONY: gen.all
gen.all: tools.verify.swagger
	@-rm -vrf $(ROOT_DIR)/internal/apiserver/dal/query
	@go run $(ROOT_DIR)/cmd/godfrey-gencode/gencode.go
	@swagger generate spec --scan-models -w $(ROOT_DIR)/cmd/godfrey-gencode -o $(ROOT_DIR)/docs/swagger.yaml
	