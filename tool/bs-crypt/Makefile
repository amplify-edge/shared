.PHONY: all

all: build-clean build test build-clean

build: build-clean
	@go build -o bs-crypt .

build-clean:
	rm -rf ./bs-crypt

test: export BS_CRYPT_PASSWORD = mXRd13vXpnw9H8i0fJbcjwLLmDHPUvrf

test:
	# e2e test
	./bs-crypt encrypt -s ./lib/encrypted -d ./lib/decrypted