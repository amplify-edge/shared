# https://github.com/go-flutter-desktop/hover

# includes
BOILERPLATE_FSPATH=./../../boilerplate

include $(BOILERPLATE_FSPATH)/help.mk
include $(BOILERPLATE_FSPATH)/tool.mk
include $(BOILERPLATE_FSPATH)/go.mk


# go
override GO_FSPATH=${PWD}

override GO_BIN_NAME=${TOOL_HOVER_BIN_NAME}
override GO_BUILD_OUT_FSPATH:=${TOOL_HOVER_BIN_FSPATH}

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

## Build
this-build: this-build-clean
	$(MAKE) go-build
	which ${TOOL_HOVER_BIN_NAME}

this-build-clean:
	rm -rf ${GO_BUILD_OUT_FSPATH}

this-test:
	${GO_BUILD_OUT_FSPATH} version
	# should get "Hover (devel)"



## Fork ( not used anymore, but left here in case we need to build a fork again to work around a problem)
LIB=github.com/go-flutter-desktop/hover
LIB_FSPATH=$(GOPATH)/src/$(LIB)
LIB_BRANCH=darwin-bundle/icns-improvement
fork-dep:
	git clone https://github.com/go-flutter-desktop/hover $(LIB_FSPATH)
	cd $(LIB_FSPATH) && git checkout $(LIB_BRANCH)
fork-dep-clean:
	rm -rf $(LIB_FSPATH)
fork-build:
	cd $(LIB_FSPATH) && go build -o $(GOPATH)/bin/hover .