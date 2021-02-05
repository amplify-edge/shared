
SHARED_FSPATH=./../../shared
BOILERPLATE_FSPATH=$(SHARED_FSPATH)/boilerplate

include $(BOILERPLATE_FSPATH)/help.mk
include $(BOILERPLATE_FSPATH)/os.mk
include $(BOILERPLATE_FSPATH)/go.mk

# ciped from shared 
include ./shared-grpc.mk

include flu-proto.mk
#include go-grpc.mk
#include go-proto.mk
include protoc.mk

# grpc tools

## All
this-all:
	$(MAKE) grpc-print
	$(MAKE) go-os-check
	$(MAKE) this-which

this-custom:
	# used for debugging
	
	$(MAKE) protoc-dep-delete
	$(MAKE) protoc-dep
	$(MAKE) protoc-which
	$(MAKE) protoc-dep-delete
	
	
## Which
this-which: flu-proto-which protoc-which

## Dep
this-dep: flu-proto-dep protoc-dep
	
## Dep delete
this-dep-delete: flu-proto-delete protoc-dep-delete

## VSCode
this-vscode: flu-proto-vscode protoc-vscode
	#code --add $(LIB_GRPCUI_REPO_FSPATH) --reuse-window
