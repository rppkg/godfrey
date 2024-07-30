.PHONY: gen.query
gen.query:
	@-rm -vrf $(ROOT_DIR)/internal/apiserver/dal/query
	@go run $(ROOT_DIR)/cmd/godfrey-gencode/gencode.go