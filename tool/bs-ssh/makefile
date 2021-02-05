# github.com/nanobox-io/golang-ssh

# includes
BOILERPLATE_FSPATH=./../../boilerplate

include $(BOILERPLATE_FSPATH)/help.mk
include $(BOILERPLATE_FSPATH)/tool.mk
include $(BOILERPLATE_FSPATH)/go.mk


# go
override GO_FSPATH=${PWD}

override GO_BIN_NAME=${TOOL_SSH_BIN_NAME}
override GO_BUILD_OUT_FSPATH:=${TOOL_SSH_BIN_FSPATH}

override GO_BUILD_OUT_ALL_FSPATH=${PWD}/assets



# Print
this-print:
	@echo
	@echo -- TOOL Hover --

	@echo GO_FSPATH 				${GO_FSPATH}
	@echo GO_BIN_NAME 				${GO_BIN_NAME}
	@echo GO_BUILD_OUT_FSPATH: 		${GO_BUILD_OUT_FSPATH}
	@echo GO_BUILD_OUT_ALL_FSPATH: 	${GO_BUILD_OUT_ALL_FSPATH}

	@echo

## Fixes go mod
this-mod:
	$(MAKE) go-mod-update

## Build
this-build: this-build-clean
	$(MAKE) go-build
	which ${GO_BIN_NAME}

this-build-clean:
	rm -f ${GO_BUILD_OUT_FSPATH}

this-test:
	${GO_BIN_NAME} -h

	${GO_BIN_NAME} version

	${GO_BIN_NAME} gen

	${GO_BIN_NAME} connect



