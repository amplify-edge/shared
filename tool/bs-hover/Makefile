# https://github.com/go-flutter-desktop/hover

# includes
SHARED_FSPATH=./../../includes

include $(SHARED_FSPATH)/help.mk

# Print
print:
	@echo
	@echo -- TOOL Hover --
	@echo

## Build
build: build-clean
	@go build  -o bs-hover .

build-clean:
	rm -rf ./bs-hover

test:
	./bs-hover version