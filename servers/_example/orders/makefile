LIB=github.com/google/go-cloud
LIB_FSPATH=$(GOPATH)/src/$(LIB)


EXAMPLE_FSPATH=$(LIB_FSPATH)/samples/order

print:
	@echo
	@echo LIB: 				$(LIB)
	@echo LIB_FSPATH: 		$(LIB_FSPATH)
	@echo EXAMPLE_FSPATH	$(EXAMPLE_FSPATH)

lib-dep:
	git clone https://$(LIB) $(LIB_FSPATH)
lib-dep-delete:
	rm -rf $(LIB_FSPATH)
lib-dep-open:
	code -add $(LIB_FSPATH)
	
run:
	#Has many mains
	cd $(EXAMPLE_FSPATH) && go run .

run-open:
	# 10538
	open http://127.0.01:10538
