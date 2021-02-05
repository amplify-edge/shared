

this-all: this-print this-dep this-build this-print-end


## Print
this-print:
	@echo
	@echo -- TOOL bs-embed : start --

	@echo bs-embed: $(EMBED_BIN_NAME)
	@echo bs-gen: $(GEN_BIN_NAME)
	@echo
	
this-print-end:
	@echo
	@echo -- TOOL bd-embed : end--
	@echo


### DEP

LIB_GOSEC=github.com/securego/gosec
LIB_GOSEC_FSPATH=$(GOPATH)/src/$(LIB_GOSEC)
LIB_GOSEC_BIN=$(GOPATH)/bin/gosec

LIB_BINDATA=github.com/go-bindata/go-bindata
LIB_BINDATA_FSPATH=$(GOPATH)/src/$(LIB_BINDATA)
LIB_BINDATA_BIN=$(GOPATH)/bin/go-bindata

this-dep: this-dep-delete

	# go get github.com/securego/gosec/v2/cmd/gosec

	# clone
	git clone https://$(LIB_GOSEC) $(LIB_GOSEC_FSPATH)
	git clone https://$(LIB_BINDATA) $(LIB_BINDATA_FSPATH)

	# build
	cd $(LIB_GOSEC_FSPATH)/cmd/gosec && go build -o $(LIB_GOSEC_BIN) .
	which $(LIB_GOSEC_BIN)
	cd $(LIB_BINDATA_FSPATH)/go-bindata && go build -o $(LIB_BINDATA_BIN) .
	which $(LIB_BINDATA_BIN)

this-dep-delete:
	# sec
	rm -rf $(LIB_GOSEC_FSPATH)
	rm -f $(LIB_GOSEC_BIN)

	# bindata
	rm -rf $(LIB_BINDATA_FSPATH)
	rm -f $(LIB_BINDATA_BIN)




### Build the Tools

EMBED_BIN_NAME=bs-embed
EMBED_BIN_FSPATH=$(GOPATH)/bin/$(EMBED_BIN_NAME)

GEN_BIN_NAME=bs-gen
GEN_BIN_FSPATH=$(GOPATH)/bin/$(GEN_BIN_NAME)

this-build: this-build-delete
	cd cmd/embed && go build -o $(EMBED_BIN_FSPATH) .
	which $(EMBED_BIN_NAME)

	cd cmd/gen && go build -o $(GEN_BIN_FSPATH) .
	which $(GEN_BIN_NAME)

this-build-delete:
	rm -f $(EMBED_BIN_NAME)
	rm -f $(GEN_BIN_NAME)

### Play

this-run:
	$(BIN_FSPATH)





