# TOOLS

# These constants are needed because the tools are used by all the other Makefiles.
# So all the other makes files can use these constants to call the tools.

# Its:
#- our golang bs binaries that always go into the gobin
#- the GRPC tools that go into specific OS global paths.


# path to binaries
TOOL_BIN_FSPATH=$(GOPATH)/bin

TOOL_DUMMY_BIN_NAME=bs-dummy
TOOL_DUMMY_BIN_FSPATH=${TOOL_BIN_FSPATH}/${TOOL_DUMMY_BIN_NAME}

TOOL_LANG_BIN_NAME=bs-lang
TOOL_LANG_BIN_FSPATH=${TOOL_BIN_FSPATH}/${TOOL_LANG_BIN_NAME}

TOOL_BOX_BIN_NAME=bs-box
TOOL_BOX_BIN_FSPATH=${TOOL_BIN_FSPATH}/${TOOL_BOX_BIN_NAME}

TOOL_HOVER_BIN_NAME=bs-hover
TOOL_HOVER_BIN_FSPATH=${TOOL_BIN_FSPATH}/${TOOL_HOVER_BIN_NAME}

TOOL_HUGO_BIN_NAME=bs-hugo
TOOL_HUGO_BIN_FSPATH=${TOOL_BIN_FSPATH}/${TOOL_HUGO_BIN_NAME}

TOOL_SSH_BIN_NAME=bs-ssh
TOOL_SSH_BIN_FSPATH=${TOOL_BIN_FSPATH}/${TOOL_SSH_BIN_NAME}

TOOL_SSHCONFIG_BIN_NAME=bs-sshconfig
TOOL_SSHCONFIG_BIN_FSPATH=${TOOL_BIN_FSPATH}/${TOOL_SSHCONFIG_BIN_NAME}

TOOL_VERSION_BIN_NAME=bs-version
TOOL_VERSION_BIN_FSPATH=${TOOL_BIN_FSPATH}/${TOOL_VERSION_BIN_NAME}

TOOL_PROTODOC_BIN_NAME=bs-protodoc
TOOL_PROTODOC_BIN_FSPATH=${TOOL_BIN_FSPATH}/${TOOL_PROTODOC_BIN_NAME}

## Print all tools
tool-print:
	@echo
	@echo -- TOOL Print: start --
	@echo

	@echo
	@echo TOOL_LIB: 									$(TOOL_LIB)
	@echo TOOL_LIB_FSPATH: 								$(TOOL_LIB_FSPATH)
	@echo
	@echo TOOL_BIN_FSPATH: 								$(TOOL_BIN_FSPATH)

	@echo
	@echo TOOL_DUMMY_BIN_NAME: 							$(TOOL_DUMMY_BIN_NAME)
	@echo TOOL_DUMMY_BIN_FSPATH: 						$(TOOL_DUMMY_BIN_FSPATH)

	@echo
	@echo TOOL_LANG_BIN_NAME: 							$(TOOL_LANG_BIN_NAME)
	@echo TOOL_LANG_BIN_FSPATH: 						$(TOOL_LANG_BIN_FSPATH)

	@echo
	@echo TOOL_HOVER_BIN_NAME: 							$(TOOL_HOVER_BIN_NAME)
	@echo TOOL_HOVER_BIN_FSPATH: 						$(TOOL_HOVER_BIN_FSPATH)

	@echo
	@echo TOOL_HUGO_BIN_NAME: 							$(TOOL_HUGO_BIN_NAME)
	@echo TOOL_HUGO_BIN_FSPATH: 						$(TOOL_HUGO_BIN_FSPATH)

	@echo
	@echo TOOL_PROTODOC_BIN_NAME: 						$(TOOL_PROTODOC_BIN_NAME)
	@echo TOOL_PROTODOC_BIN_FSPATH: 					$(TOOL_PROTODOC_BIN_FSPATH)
	
	@echo
	@echo -- TOOL Print : end --
	@echo

