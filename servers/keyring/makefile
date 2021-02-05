
LIB=Google Cloud deploy
LIB_FSPATH=$(GOPATH)/src/$(LIB)

BIN_FSPATH=$(PWD)/_bin
LIB_BIN=$(BIN_FSPATH)/keyring

print:
	@echo
	@echo LIB: 				$(LIB)
	@echo LIB_FSPATH: 		$(LIB_FSPATH)
	@echo LIB_BIN: 			$(LIB_BIN)

lib-dep:
	git clone https://$(LIB) $(LIB_FSPATH)
lib-dep-delete:
	rm -rf $(LIB_FSPATH)
lib-dep-open:
	code -add $(LIB_FSPATH)

build:
	cd $(LIB_FSPATH)/cmd/keyring && go build -o $(LIB_BIN) .

SERVICE_NAME=srv01
USER_NAME=joe
SECRET_VALUE=secretpassword

run-help:

run-create:
	# set value
	echo -n "$(SECRET_VALUE)" | $(LIB_BIN) set  $(SERVICE_NAME) $(USER_NAME) --yes

run-all:
	# via stdin set, get, delete and get exercises everything as a basic test it works.
	
	# set value
	echo -n "$(SECRET_VALUE)" | $(LIB_BIN) set  $(SERVICE_NAME) $(USER_NAME) --yes

	# get value
	$(LIB_BIN) get $(SERVICE_NAME) $(USER_NAME)

	# delete value
	$(LIB_BIN) delete $(SERVICE_NAME) $(USER_NAME)

	# get value ( should fail )
	$(LIB_BIN) get $(SERVICE_NAME) $(USER_NAME)

