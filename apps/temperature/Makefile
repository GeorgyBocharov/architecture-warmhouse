BIN := "./bin/temperature-app"
DOCKER_IMG := "temperature-app:main"
SERVER_BASE_DIR := server
OAPI_CODEGEN_VERSION := v2.5.1


build:
	go build -v -o $(BIN) ./main

build-img:
	docker build \
		-t $(DOCKER_IMG) \
		-f Dockerfile .

generate: install-generate-deps generate-server
	go generate ./...

install-generate-deps: 
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@$(OAPI_CODEGEN_VERSION)
	
generate-server: 
	oapi-codegen -config server/cfg.yaml -package server  ./api/openapi.yaml