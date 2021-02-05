
BIN_PATH=$(PWD)/_bin
DATA_PATH=$(PWD)/_data

MINIO_CLI_BIN=$(BIN_PATH)/minio-cli
MINIO_SERVER_BIN=$(BIN_PATH)/minio-server
MINIO_SERVER_DATA=$(DATA_PATH)/minio-server

MINI_CLI_CONFIG_PATH=/Users/apple/.minio-cli
MINIO_ALIAS=myminio
MINIO_ACCESS_KEY=minioadmin
MINIO_SECRET_KEY=minioadmin

print:
	@echo
	@echo BIN_PATH: 			$(BIN_PATH)
	@echo
	@echo MINIO_CLI_BIN: 		$(MINIO_CLI_BIN)
	@echo MINIO_SERVER_BIN: 	$(MINIO_SERVER_BIN)
	@echo MINIO_SERVER_DATA: 	$(MINIO_SERVER_DATA)
	@echo

### DEPS

dep-os-minio-server:
	# https://min.io/download#/macos
	# minio server
	curl -O https://dl.min.io/server/minio/release/darwin-amd64/minio
	cp ./minio $(MINIO_SERVER_BIN)
	rm -rf ./minio
	chmod +x $(MINIO_SERVER_BIN)
dep-os-minio-cli:
	# minio client
	curl -O https://dl.min.io/client/mc/release/darwin-amd64/mc
	cp ./mc $(MINIO_CLI_BIN)
	rm -rf ./mc
	chmod +x $(MINIO_CLI_BIN)
dep-os-minio-delete:
	rm -rf $(MINIO_CLI_BIN)
	rm -rf $(MINIO_SERVER_BIN)


### RUN

run-minio-server:
	#  http://192.168.2.100:9000  http://127.0.0.1:9000  
	MINIO_ACCESS_KEY=$(MINIO_ACCESS_KEY) MINIO_SECRET_KEY=$(MINIO_SECRET_KEY) $(MINIO_SERVER_BIN) server $(MINIO_SERVER_DATA)
run-minio-client-init:
	$(MINIO_CLI_BIN) --help

	# https://docs.min.io/docs/minio-client-quickstart-guide

	# list from their test server that is preconfigured.
	$(MINIO_CLI_BIN) ls play

	# setup config to local minio server
	$(MINIO_CLI_BIN) alias set $(MINIO_ALIAS) http://127.0.0.1:9000 $(MINIO_ACCESS_KEY) $(MINIO_SECRET_KEY)

run-minio-client-watch:
	# Can see changes
	$(MINIO_CLI_BIN) watch $(MINIO_ALIAS)

run-minio-client-upload:
	# copy a file up
	$(MINIO_CLI_BIN) mb --ignore-existing $(MINIO_ALIAS)/mybucket
	$(MINIO_CLI_BIN) cp test01.md $(MINIO_ALIAS)/mybucket

	# list my buckets
	$(MINIO_CLI_BIN) tree --files $(MINIO_ALIAS)

	# list my files
	$(MINIO_CLI_BIN) ls --recursive $(MINIO_ALIAS)
	
run-minio-client-update:
	# mc admin update $(MINIO_ALIAS)
	$(MINIO_CLI_BIN) admin update $(MINIO_ALIAS)
