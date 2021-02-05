


## All
this-all: this-print this-build this-print-end


## Print
this-print:
	@echo
	@echo -- TOOL bs-embed-ex: start --
	@echo
	
this-print-end:
	@echo
	@echo -- TOOL bd-embed-ex : end--
	@echo


NAME=maintemplate
MOD_LIST= maintemplate mod01
EMBED_FLAGS = -pkg assets -nomemcopy -fs -prefix "assets"


this-build: this-build-delete

	# Generate phase
	@echo 
	go generate ./...
	@echo "[OK] Generate phase - Files added to embed "
	@echo 

	# Embed Phase
	# ensure you have a .gitignore to avoid commiting bindata.go files into git.
	@echo 
	for mod in $(MOD_LIST); do \
  		go-bindata -o ./$$mod/assets/bindata.go ${EMBED_FLAGS} ./$$mod/static/...; \
  	done

	@echo "[OK] Embed phase - Files embedded to bindata.go "
	@echo 

	# security check
	@echo 
	@gosec -no-fail ./...
	@echo "[OK] Go security check was completed!"
	@echo 

	# Build phase
	@echo 
	@go build -o ./build/$(NAME) ./maintemplate/cmd/app/*.go
	@echo "[OK] App binary was created!"
	@echo 


this-build-delete:
	rm -f /build/$(NAME)

this-run:
	@./build/$(NAME)