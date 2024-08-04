.PHONY: doc.serve
doc.serve: tools.verify.swagger
	@swagger serve -F=redoc --no-open --port 18087 $(ROOT_DIR)/docs/swagger.yaml