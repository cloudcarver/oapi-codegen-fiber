SHELL := /bin/zsh
PROJECT_DIR=$(shell pwd)

OAPI_CODEGEN_VERSION=v2.4.1
OAPI_CODEGEN_BIN=$(PROJECT_DIR)/bin/oapi-codegen
OAPI_GEN_DIR=$(PROJECT_DIR)/internal/apigen
OAPI_CODEGEN_FIBER_BIN=$(PROJECT_DIR)/bin/oapi-codegen-fiber

install-oapi-codegen:
	@DIR=$(PROJECT_DIR)/bin VERSION=${OAPI_CODEGEN_VERSION} ./scripts/install-oapi-codegen.sh

gen: install-oapi-codegen
	go run main.go --path test/v1.yaml --out test/out.go --package apigen
	${OAPI_CODEGEN_BIN} -generate=fiber,client,types -package apigen -o test/api.gen.go test/v1.yaml
