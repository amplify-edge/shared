# protoc

## Protoc which
protoc-which:
	@echo
	@echo -- protoc : which --
	@which protoc
	protoc --version
	@echo

## Protoc dep
protoc-dep:
	# http://google.github.io/proto-lens/installing-protoc.html

	@echo
	@echo -- protoc : dep--

ifeq ($(GO_OS), windows)
	@echo Windows detected

else
	
ifeq ($(GO_OS), linux)
	@echo Linux detected

else
	@echo Darwin detected
	curl -OL $(LIB_PROTOC_URL)
	sudo unzip -o $(LIB_PROTOC_FILENAME_DARWIN) -d /usr/local bin/protoc
	sudo unzip -o $(LIB_PROTOC_FILENAME_DARWIN) -d /usr/local 'include/*'
	rm -f $(LIB_PROTOC_FILENAME_DARWIN)
endif
endif

## Protoc-git-delete
protoc-dep-delete:

	@echo
	@echo -- protoc : dep-delete --

ifeq ($(GO_OS), windows)
	@echo Windows detected

else
	
ifeq ($(GO_OS), linux)
	@echo Linux detected

else
	@echo Darwin detected
	#brew uninstall protobuf
	rm -f /usr/local/bin/protoc
endif
endif

protoc-vscode:
	# NONE

