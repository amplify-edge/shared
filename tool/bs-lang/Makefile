# Don't use too many makefiles variables it's confusing.
BOILERPLATE_FSPATH=../../includes

include $(BOILERPLATE_FSPATH)/help.mk

# Print
print:
	@echo
	@echo -- TOOL LANG --
	@echo

## Build
build: build-clean
	go build -o bs-lang .

build-clean:
	rm -rf ./bs-lang

test:
	go test -v ./services

# lang tool usage. TODO
# 1. make build ==> build lang tool
# 2. make lang-full-json-flutter ==> generate translated json files
# 3. make lang-arb-flutter ==> convert json to arb

## Example json
example-json-flutter:
	# generate full json files
	## DOES NOT WORK !!
	${TOOL_LANG_BIN_NAME} flutter --dir examples/lang_example/lib/lang/generated/ --template examples/lang_example/lib/lang/generated/intl_messages.arb --prefix mockText --languages en,fr,es,ur -f

## Example arb
example-arb-flutter:
	# ??
	## DOES NOT WORK !!!
	${TOOL_LANG_BIN_NAME} flutter --dir examples/lang_example/lib/lang/generated/
