
#BOILERPLATE_FSPATH=./../boot/boilerplate
BOILERPLATE_FSPATH=./boilerplate

include $(BOILERPLATE_FSPATH)/help.mk
include $(BOILERPLATE_FSPATH)/os.mk
include $(BOILERPLATE_FSPATH)/gitr.mk
include $(BOILERPLATE_FSPATH)/tool.mk
include $(BOILERPLATE_FSPATH)/flu.mk
include $(BOILERPLATE_FSPATH)/go.mk

# remove the "v" prefix
VERSION ?= $(shell echo $(TAGGED_VERSION) | cut -c 2-)

## Build in CI
this-all: this-print this-dep this-build this-print-end

this-print:
	@echo
	@echo -- SHARED REPO : start --
	@echo

this-print-end:
	@echo
	@echo -- SHARED REPO ; end --
	@echo

## Print all settings
this-print-all: ## print
	
	$(MAKE) os-print
	
	$(MAKE) gitr-print

	$(MAKE) go-print

	$(MAKE) tool-print
	
	$(MAKE) flu-print

	$(MAKE) flu-gen-lang-print

this-ci-check:
	# just to check calling make from CI works

	@echo "Hello CI"


this-dep:
	# none
	cd ./dep && $(MAKE) this-all


this-build:
	# build
	cd ./tool && $(MAKE) this-all

	# SDK
	#cd ./sdk && $(MAKE) this-all


### Sync

# From the Shared repo, we copy the boilerplate out to the others repos
# This is for devs to have a singel make function 

# Repos

REPO_LIST=main mod sys sys-share dev


# Folders in each repo
CI_FOLDER_SOURCE_NAME=./ci-templates/workflows
CI_FOLDER_TARGET_NAME=./.github


#override GITR_COMMIT_MESSAGE = joe
override GITR_COMMIT_MESSAGE = $(M)

this-git-all: this-copy-all this-git-commit-all this-git-catchup-all this-git-push-all


## Copy boilerplate to other repos.
this-copy-all:

	for repo in $(REPO_LIST); do \
		cp -Rvi $(CI_FOLDER_SOURCE_NAME) ./../$$repo/$(CI_FOLDER_TARGET_NAME) ; \
  	done
	  
## Forces commit in all repos
this-git-commit-all: 

	# Useful when working across many repos.
	# Add the same Issue number

	@echo GITR_COMMIT_MESSAGE: $(GITR_COMMIT_MESSAGE)

	for repo in $(REPO_LIST); do \
		cd ./../$$repo && $(MAKE) gitr-fork-commit ; \
  	done

## Force catchup from Upsteam for all repos.
this-git-catchup-all:

	for repo in $(REPO_LIST); do \
		cd ./../$$repo && $(MAKE) gitr-fork-catchup ; \
  	done

this-git-push-all:

	for repo in $(REPO_LIST); do \
		cd ./../$$repo && $(MAKE) gitr-fork-push ; \
  	done
