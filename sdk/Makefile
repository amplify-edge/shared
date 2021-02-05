

## Build all
this-all: this-print this-tool-build this-build this-print-end
	

## Print all settings
this-print: 
	@echo
	@echo "-- SDK : start --"
	@echo

this-print-end:
	@echo
	@echo "-- SDK : end --"
	@echo
	@echo
	
this-build:
	go build .
	
this-tool-build:
	cd bs-embed && $(MAKE) this-all

this-ex-build:
	cd bs-embed-ex && $(MAKE) this-all

this-ex-run:
	cd bs-embed-ex && $(MAKE) this-run
