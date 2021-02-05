
BIN_PATH=$(PWD)/_bin
DATA_PATH=$(PWD)/_data

NATS_CLI_BIN=$(BIN_PATH)/nats-cli
NATS_SERVER_BIN=$(BIN_PATH)/nats-server

print:
	@echo
	@echo BIN_PATH: 			$(BIN_PATH)
	@echo
	@echo NATS_CLI_BIN: 		$(NATS_CLI_BIN)
	@echo NATS_SERVER_BIN: 		$(NATS_SERVER_BIN)
	@echo

### DEPS

dep-os-nats-server:
	# https://docs.nats.io/nats-server/installation
	curl -L https://github.com/nats-io/nats-server/releases/download/v2.0.0/nats-server-v2.0.0-darwin-amd64.zip -o nats-server.zip
	unzip nats-server.zip -d nats-server
	rm -rf ./nats-server.zip
	cp ./nats-server/nats-server-v2.0.0-darwin-amd64/nats-server $(NATS_SERVER_BIN)
	rm -rf ./nats-server
dep-os-nats-cli:
	curl -L https://github.com/nats-io/natscli/releases/download/0.0.21/nats-0.0.21-darwin-amd64.zip -o nats-cli.zip
	unzip nats-cli.zip -d nats-cli
	rm -rf ./nats-cli.zip
	cp ./nats-cli/nats-0.0.21-darwin-amd64/nats $(NATS_CLI_BIN)
	rm -rf ./nats-cli
dep-os-nats-delete:
	rm -rf $(NATS_SERVER_BIN)
	rm -rf $(NATS_CLI_BIN)

### RUN
run-nats-server:
	# 0.0.0.0:4222
	$(NATS_SERVER_BIN)
run-nats-cli:
	# 0.0.0.0:4222
	$(NATS_CLI_BIN)
