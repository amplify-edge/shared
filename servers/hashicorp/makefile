
BIN_PATH=$(PWD)/_bin
DATA_PATH=$(PWD)/_data

GO_OS ?= $(shell go env GOOS)
GO_ARCH ?= $(shell go env GOARCH)

# TODO : Fork this ...
LIB_YURT_NAME=yurt
LIB_YURT=github.com/ncabatoff/$(LIB_YURT_NAME)
LIB_YURT_FSPATH=$(GOPATH)/src/$(LIB_YURT)

BIN_YURT_GET_BINARY=$(BIN_PATH)/yurt-get-binary
BIN_YURT_SINGLE=$(BIN_PATH)/yurt-run
BIN_YURT_CLUSTER=$(BIN_PATH)/yurt-cluster



print:
	@echo
	@echo BIN_PATH: 			$(BIN_PATH)
	@echo DATA_PATH: 			$(DATA_PATH)
	@echo
	@echo GO_OS: 				$(GO_OS)
	@echo GO_ARCH: 				$(GO_ARCH)
	@echo
	@echo LIB_YURT: 			$(LIB_YURT)
	@echo LIB_YURT_FSPATH: 		$(LIB_YURT_FSPATH)
	@echo
	@echo BIN_YURT_GET_BINARY: 	$(BIN_YURT_GET_BINARY)
	@echo BIN_YURT_SINGLE: 		$(BIN_YURT_SINGLE)
	@echo BIN_YURT_CLUSTER: 	$(BIN_YURT_CLUSTER)

### DEPS

dep-os:
	# .env 
	go get github.com/joho/godotenv/cmd/godotenv

dep-lib:
	git clone https://$(LIB_YURT) $(LIB_YURT_FSPATH)

dep-lib-delete:
	rm -rf $(LIB_YURT_FSPATH)

dep-open:
	# ad to workspace
	code --add $(LIB_YURT_FSPATH)

build:
	cd $(LIB_YURT_FSPATH)/cmd/get-binary && go build -o $(BIN_YURT_GET_BINARY) 
	cd $(LIB_YURT_FSPATH)/cmd/yurt-run && go build -o $(BIN_YURT_SINGLE)
	cd $(LIB_YURT_FSPATH)/cmd/yurt-cluster && go build -o $(BIN_YURT_CLUSTER)


run-yurt-get-binary:
	# Gets the binaries
	$(BIN_YURT_GET_BINARY) -arch=$(GO_ARCH) -os=$(GO_OS) -version=1.7.3 -workdir=$(DATA_PATH)

run-yurt-single:
	# Connects to your running cluster
	cp yurt-run.yml $(BIN_PATH)
	$(BIN_YURT_SINGLE) -config-file=yurt-run.yml

run-yurt-cluster:
	# Runs the Cluster
	# Defaults
	# Vault GUI: http://127.0.0.1:23001/
	# Consul GUI: http://127.0.0.1:23010/
	# Nomad GUI:  http://127.0.0.1:23051/
	$(BIN_YURT_CLUSTER)
