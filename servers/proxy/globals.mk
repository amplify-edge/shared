# basd on https://theterminalguy.sh/posts/build-an-ngrok-alternative-with-inlets-and-caddy/

# Stack
	# DNS
		# demo.mydoman.com via a DNS that supports ACME
	# Public Server
		# Inlets Server exposing 443
	# Private Server
		# Caddy exposing 80
		# Any golang Service exposing 8080

GLOBAL_BIN_PATH=$(PWD)/_bin
GLOBAL_DATA_PATH=$(PWD)/_data

globals-print:
	@echo
	@echo GLOBAL_BIN_PATH: 			$(GLOBAL_BIN_PATH)
	@echo GLOBAL_DATA_PATH: 		$(GLOBAL_DATA_PATH)
	@echo
