# go-proto

BOILERPLATE_FSPATH=./../boilerplate

include $(BOILERPLATE_FSPATH)/help.mk
include $(BOILERPLATE_FSPATH)/os.mk
include $(BOILERPLATE_FSPATH)/grpc.mk

## Protoc which
go-proto-which:
	@echo
	@echo -- go-proto : which --
	@which go-proto
	go-proto --version
	@echo

## Protoc dep
go-proto-dep:
	# http://google.github.io/proto-lens/installing-go-proto.html

	@echo
	@echo -- go-proto : dep--

ifeq ($(GO_OS), windows)
	@echo Windows detected

else
	
ifeq ($(GO_OS), linux)
	@echo Linux detected

else
	@echo Darwin detected
	brew install protobuf@$(LIB_PROTOC_VERSION)
endif
endif

## Protoc-git-delete
go-proto-dep-delete:

	@echo
	@echo -- go-proto : dep-delete --

ifeq ($(GO_OS), windows)
	@echo Windows detected

else
	
ifeq ($(GO_OS), linux)
	@echo Linux detected

else
	@echo Darwin detected
	brew uninstall protobuf
endif
endif

go-proto-vscode-add:
	# NONE

## All-build
go-proto-build: 
	# NONE

