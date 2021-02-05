# basd on https://theterminalguy.sh/posts/build-an-ngrok-alternative-with-inlets-and-caddy/

# Stack
	# DNS
		# demo.mydoman.com via a DNS that supports ACME
	# Public Server
		# Inlets Server exposing 443
	# Private Server
		# Caddy exposing 80
		# Any golang Service exposing 8080

INLETS_CLI_BIN=$(GLOBAL_BIN_PATH)/inlets-cli
INLETS_SERVER_BIN=$(GLOBAL_BIN_PATH)/inlets-server
INLETS_SERVER_DATA=$(GLOBAL_DATA_PATH)/inlets-server

INLETS_CLI_CONFIG_PATH=/Users/apple/.minio-cli
INLETS_ALIAS=myminio
INLETS_ACCESS_KEY=minioadmin
INLETS_SECRET_KEY=minioadmin

inlets-print:
	@echo
	@echo INLETS_CLI_BIN: 		$(INLETS_CLI_BIN)
	@echo INLETS_SERVER_BIN: 	$(INLETS_SERVER_BIN)
	@echo INLETS_SERVER_DATA: 	$(INLETS_SERVER_DATA)
	@echo

### DEPS

inlets-dep-init:
	mkdir -p $(GLOBAL_BIN_PATH)
inlets-dep-delete:
	rm -rf $(INLETS_CLI_BIN)
	rm -rf $(INLETS_SERVER_BIN)

inlets-dep: inlets-dep-server inlets-dep-cli
	
inlets-dep-server: inlets-dep-init
	# inlets server
	curl -L  https://github.com/inlets/inlets/releases/download/2.7.10/inlets-darwin -o inlets-darwin
	cp ./inlets-darwin $(INLETS_SERVER_BIN)
	rm -rf ./inlets-darwin
	chmod +x $(INLETS_SERVER_BIN)

inlets-dep-cli: inlets-dep-init
	# inlets client
	curl -L  https://github.com/inlets/inletsctl/releases/download/0.8.0/inletsctl-darwin -o inletsctl-darwin
	cp ./inletsctl-darwin $(INLETS_CLI_BIN)
	rm -rf ./inletsctl-darwin
	chmod +x $(INLETS_CLI_BIN)

### RUN

inlets-run-server:
	$(INLETS_SERVER_BIN) -h

inlets-run-cli:
	$(INLETS_CLI_BIN) -h