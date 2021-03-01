# https://github.com/invertase/melos

# Examples:
# https://github.com/fluttercommunity/plus_plugins


BIN_PATH=$(PWD)/_bin
DATA_PATH=$(PWD)/_data

LIB_MELOS_NAME=melos
LIB_MELOS=github.com/invertase/$(LIB_MELOS_NAME)

LIB_MELOS_PLUS_NAME=plus_plugins
LIB_MELOS_PLUS=github.com/fluttercommunity/$(LIB_MELOS_PLUS_NAME)

LIB_MELOS_NEW=amp_new
LIB_MELOS_NEW_PKG01=amp_new_01
LIB_MELOS_NEW_PKG02=amp_new_02



print:
	@echo
	@echo BIN_PATH: 			$(BIN_PATH)
	@echo DATA_PATH: 			$(DATA_PATH)
	@echo
	@echo LIB_MELOS: 			$(LIB_MELOS)
	@echo LIB_MELOS_SCENES: 	$(LIB_MELOS_PLUS)


	@echo LIB_MELOS_NEW: 		$(LIB_MELOS_NEW)
	@echo LIB_MELOS_NEW_PKG01: 	$(LIB_MELOS_NEW_PKG01)
	@echo LIB_MELOS_NEW_PKG02: 	$(LIB_MELOS_NEW_PKG02)

	@echo

### DEPS

dep:
	git clone https://$(LIB_MELOS)
	git clone https://$(LIB_MELOS_PLUS)

dep-delete:
	rm -rf $(LIB_MELOS_NAME)
	rm -rf $(LIB_MELOS_PLUS_NAME)


### RUN
melos-dep:
	dart pub global activate melos
melos-dep-delete:
	dart pub global deactivate melos

### Plus exisiting example

run-plus:
	# Shows working with melow on a repo that heavily uses it.
	cd $(LIB_MELOS_PLUS_NAME) && melos --help

	cd $(LIB_MELOS_PLUS_NAME) && melos version

	cd $(LIB_MELOS_PLUS_NAME) && melos list


### New example

run-new-create:
	# Show working with melos on a new repo.
	mkdir -p $(LIB_MELOS_NEW)
	cd $(LIB_MELOS_NEW) && flutter create $(LIB_MELOS_NEW_PKG01)
	cd $(LIB_MELOS_NEW) && flutter create $(LIB_MELOS_NEW_PKG02)
run-new-delete:
	rm -rf $(LIB_MELOS_NEW)

run-new-melos-init:
	cd $(LIB_MELOS_NEW) && melos init