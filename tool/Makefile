# Don't use too many makefiles variables it's confusing.
BOILERPLATE_FSPATH=../includes

include $(BOILERPLATE_FSPATH)/help.mk

all: print build test

## Print all settings
print:
	@echo tools

## Build all tools
build:
	@echo -- TOOLS - BUILD ALL: start --

	@echo Dummy
	cd ./bs-dummy && $(MAKE) build

	@echo Lang
	cd ./bs-lang && $(MAKE) build

	@echo Hover
	cd ./bs-hover && $(MAKE) build

	@echo Hugo
	cd ./bs-hugo && $(MAKE) build

	@echo -- TOOLS - BUILD ALL: finish --

## Clean the build
build-clean:
	@echo -- TOOLS - CLEAN ALL: start --

	@echo Dummy
	cd ./bs-dummy && $(MAKE) build-clean

	@echo Lang
	cd ./bs-lang && $(MAKE) build-clean

	@echo Hover
	cd ./bs-hover && $(MAKE) build-clean

	@echo Hugo
	cd ./bs-hugo && $(MAKE) build-clean

	# -- PLUS ALL OTHER TOOLS !!

	@echo -- TOOLS - CLEAN ALL: finish --

## Test all tools
test:
	@echo -- TOOLS - TEST ALL: start --

	@echo Dummy
	cd ./bs-dummy && $(MAKE) test

	@echo Lang
	cd ./bs-lang && $(MAKE) test

	@echo Hover
	cd ./bs-hover && $(MAKE) test

	@echo Hugo
	cd ./bs-hugo && $(MAKE) test

	@echo -- TOOLS - TEST ALL: finish --


