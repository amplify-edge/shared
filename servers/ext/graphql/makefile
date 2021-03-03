# https://github.com/gql-dart/gql

# Examples:
# https://github.com/gql-dart/ferry
# 


BIN_PATH=$(PWD)/_bin
DATA_PATH=$(PWD)/_data

LIB_GQL_NAME=gql
LIB_GQL=github.com/gql-dart/$(LIB_GQL_NAME)

LIB_FERRY_NAME=ferry
LIB_FERRY=github.com/gql-dart/$(LIB_FERRY_NAME)


print:
	@echo
	@echo BIN_PATH: 			$(BIN_PATH)
	@echo DATA_PATH: 			$(DATA_PATH)
	@echo
	@echo LIB_GQL: 			$(LIB_GQL)
	@echo LIB_FERRY: 		$(LIB_FERRY)
	
	@echo

### DEPS

dep:
	git clone https://$(LIB_GQL)
	git clone https://$(LIB_FERRY)
	

dep-delete:
	rm -rf $(LIB_GQL_NAME)
	rm -rf $(LIB_FERRY_NAME)
	rm -rf $(LIB_GQL_SCENES_NAME)


### RUN
run-doc:
	cd $(LIB_FERRY_NAME) && flutter run -d chrome
run-scenes:
	cd $(LIB_GQL_SCENES_NAME) && flutter run -d chrome
