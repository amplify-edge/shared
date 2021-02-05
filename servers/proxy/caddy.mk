
CADDY_SERVER_BIN=$(GLOBAL_BIN_PATH)/caddy-server
CADDY_SERVER_DATA=$(GLOBAL_DATA_PATH)/caddy-server

CADDY_CLI_CONFIG_PATH=/Users/apple/.minio-cli
CADDY_ALIAS=myminio
CADDY_ACCESS_KEY=minioadmin
CADDY_SECRET_KEY=minioadmin

caddy-print:
	@echo
	@echo CADDY_SERVER_BIN: 		$(CADDY_SERVER_BIN)
	@echo CADDY_SERVER_DATA: 		$(CADDY_SERVER_DATA)
	@echo

### DEPS

caddy-dep-init:
	mkdir -p $(GLOBAL_BIN_PATH)
caddy-dep-delete:
	rm -rf $(CADDY_CLI_BIN)
	rm -rf $(CADDY_SERVER_BIN)

caddy-dep: caddy-dep-init
	# caddy server
	curl -L https://github.com/caddyserver/caddy/releases/download/v2.3.0/caddy_2.3.0_mac_amd64.tar.gz -o caddy.tar.gz
	tar -xzf rcaddy.tar.gz
	#cp ./caddy-darwin $(CADDY_SERVER_BIN)
	#rm -rf ./caddy-darwin
	#chmod +x $(CADDY_SERVER_BIN)

### RUN

caddy-run-server:
	$(CADDY_SERVER_BIN) -h
