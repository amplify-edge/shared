# go-grpc

BOILERPLATE_FSPATH=./../boilerplate

include $(BOILERPLATE_FSPATH)/help.mk
include $(BOILERPLATE_FSPATH)/os.mk
include $(BOILERPLATE_FSPATH)/grpc.mk

## Protoc which
go-grpc-which:
	@echo
	@echo -- go-grpc : which --
	@which go-grpc
	go-grpc --version
	@echo

## Protoc dep
go-grpc-dep:
	# http://google.github.io/proto-lens/installing-go-grpc.html

	@echo
	@echo -- go-grpc : dep--

ifeq ($(GO_OS), windows)
	@echo Windows detected

else
	
ifeq ($(GO_OS), linux)
	@echo Linux detected

else
	@echo Darwin detected
	#
endif
endif

## Protoc-git-delete
go-grpc-dep-delete:

	@echo
	@echo -- go-grpc : dep-delete --

ifeq ($(GO_OS), windows)
	@echo Windows detected

else
	
ifeq ($(GO_OS), linux)
	@echo Linux detected

else
	@echo Darwin detected
	#brew uninstall protobuf
endif
endif

go-grpc-vscode:
	# NONE

