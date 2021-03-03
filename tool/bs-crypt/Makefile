.PHONY: all

all:

build: build-clean
	@go build -o bs-crypt .

build-clean:
	rm -rf ./bs-crypt

test:
	$(shell export BS_CRYPT_PASSWORD=blehbleh2341298astk)
# TODO
#	go test -v ./lib