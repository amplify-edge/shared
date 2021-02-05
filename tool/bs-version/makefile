# github.com/nanobox-io/golang-ssh

# includes
BOILERPLATE_FSPATH=./../../boilerplate

include $(BOILERPLATE_FSPATH)/help.mk
include $(BOILERPLATE_FSPATH)/tool.mk
include $(BOILERPLATE_FSPATH)/go.mk


# go
override GO_FSPATH=${PWD}

override GO_BIN_NAME=${TOOL_VERSION_BIN_NAME}
override GO_BUILD_OUT_FSPATH:=${TOOL_VERSION_BIN_FSPATH}

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

### Build

this-build: go-build-clean
	$(MAKE) go-build
	which ${GO_BIN_NAME}

this-build-all: go-build-all-clean
	$(MAKE) go-build-all
	ls ${GO_BUILD_OUT_ALL_FSPATH} 



### RUN


this-run-backup:
	# VERY IMPORTANT - Run this BEFORE you run any tests, so that your local .ssh folder is backedup
	@echo -n "Are you sure? [y/N] " && read ans && [ $${ans:-N} = y ]

	echo ${NOW}
	cp -r ~/.ssh/ ~/.sshbackup-$(NOW)
	code ~/.sshbackup-$(NOW)


this-run-init:
	# Then create the files that describes your ssh aspects.
	touch ${SSH_CONFIG_FSPATH}

this-run-00:
	cp ./example/00/assh.yml ${SSH_CONFIG_FSPATH}

	${MAKE} this-run-e2e

	${MAKE} this-run-open

this-run-01:
	cp ./example/01/assh.yml ${SSH_CONFIG_FSPATH}

	${MAKE} this-run-e2e

	${MAKE} this-run-open

this-run-e2e:
	# runs an E2E test that in idempotent.

	@echo
	@echo -- help --
	${GO_BIN_NAME} -h

	@echo
	@echo -- version --
	${GO_BIN_NAME} -v

	@echo
	@echo -- get --
	${GO_BIN_NAME} get

	@echo
	@echo -- patch --
	${GO_BIN_NAME} patch


this-run-open:
	# See what it does under the hood
	code ~/.ssh/assh.yml
	code ~/.ssh/config

