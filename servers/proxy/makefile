# basd on https://theterminalguy.sh/posts/build-an-ngrok-alternative-with-inlets-and-caddy/

# Stack
	# DNS
		# demo.mydoman.com via a DNS that supports ACME
	# Public Server
		# Inlets Server exposing 443
	# Private Server
		# Caddy exposing 80
		# Any golang Service exposing 8080

include ./globals.mk
include ./inlets.mk
include ./caddy.mk


INLETS_CLI_CONFIG_PATH=/Users/apple/.minio-cli
INLETS_ALIAS=myminio
INLETS_ACCESS_KEY=minioadmin
INLETS_SECRET_KEY=minioadmin

print:
	@echo
	$(MAKE) globals-print
	@echo
	$(MAKE) inlets-print
	@echo
	$(MAKE) caddy-print
	@echo

### DEPS

dep:
	$(MAKE) inlets-dep
	$(MAKE) caddy-dep
dep-delete:
	$(MAKE) inlets-dep-delete
	$(MAKE) caddy-dep-delete

### RUN

run-inlets-server:
	$(INLETS_SERVER_BIN) -h

run-inlets-cli:
	$(INLETS_CLI_BIN) -h