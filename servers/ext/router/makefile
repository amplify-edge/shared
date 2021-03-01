# https://github.com/lulupointu/vrouter

# Examples:
# https://github.com/lulupointu/vrouter_website
# https://github.com/lulupointu/vrouter_navigator_scenarios


BIN_PATH=$(PWD)/_bin
DATA_PATH=$(PWD)/_data

LIB_VROUTER_NAME=vrouter
LIB_VROUTER=github.com/lulupointu/$(LIB_VROUTER_NAME)

LIB_VROUTER_DOC_NAME=vrouter_website
LIB_VROUTER_DOC=github.com/lulupointu/$(LIB_VROUTER_DOC_NAME)

LIB_VROUTER_SCENES_NAME=vrouter_navigator_scenarios
LIB_VROUTER_SCENES=github.com/lulupointu/$(LIB_VROUTER_SCENES_NAME)

print:
	@echo
	@echo BIN_PATH: 			$(BIN_PATH)
	@echo DATA_PATH: 			$(DATA_PATH)
	@echo
	@echo LIB_VROUTER: 			$(LIB_VROUTER)
	@echo LIB_VROUTER_DOC: 		$(LIB_VROUTER_DOC)
	@echo LIB_VROUTER_SCENES: 	$(LIB_VROUTER_SCENES)
	@echo

### DEPS

dep:
	git clone https://$(LIB_VROUTER)
	git clone https://$(LIB_VROUTER_DOC)
	git clone https://$(LIB_VROUTER_SCENES)

dep-delete:
	rm -rf $(LIB_VROUTER_NAME)
	rm -rf $(LIB_VROUTER_DOC_NAME)
	rm -rf $(LIB_VROUTER_SCENES_NAME)


### RUN
run-doc:
	cd $(LIB_VROUTER_DOC_NAME) && flutter run -d chrome
run-scenes:
	cd $(LIB_VROUTER_SCENES_NAME) && flutter run -d chrome
