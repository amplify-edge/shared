
# git general stuff

# hardcoded
GITR_SERVER ?= github.com
GITR_ORG_UPSTREAM ?= amplify-edge

# A dev will fork via github, git clone, and then we need to add the upstream for them by reflecting off git origin data in the .git/config

# # calculate the upstream from the Git config.
GITR_ORG_ORIGIN=$(shell basename $(dir $(abspath $(dir $$PWD))))
#GITR_USER=$(GITR_ORG_ORIGIN)
GITR_USER=$(shell git config user.name)		# Maybe move this to a gitr-env.mk that a user has gitignored and can then override thngs.
GITR_REPO_NAME=$(notdir $(shell pwd))

# calculated
# upstream
GITR_REPO_UPSTREAM_URL=https:///$(GITR_SERVER)/$(GITR_ORG_UPSTREAM)/$(GITR_REPO_NAME)

GITR_REPO_ORIGIN_URL=https:///$(GITR_SERVER)/$(GITR_ORG_ORIGIN)/$(GITR_REPO_NAME)
GITR_REPO_ORIGIN_FSPATH=$(GOPATH)/src/$(GITR_SERVER)/$(GITR_ORG_ORIGIN)/$(GITR_REPO_NAME)

# tags and versions
GITR_LAST_TAG=$(shell git describe --exact-match --tags $(shell git rev-parse HEAD))
GITR_VERSION ?= $(shell echo $(TAGGED_VERSION) | cut -c 2-) # remove the "v" prefix

# commit message setup to be overriden
GITR_COMMIT_MESSAGE ?= autocommit


#GITR_BRANCH_NAME=main # Later for gitea
GITR_BRANCH_NAME=master

gitr-print-raw:
	# reflected from the .git/config itself
	@echo
	@echo user.name: $(shell git config user.name)
	@echo user.email: $(shell git config user.email)
	@echo
	@echo SHOULD remote.origin.url=git@github.com-gerardwebb:gerardwebb/amp-shared
	@echo remote.origin.url:  $(shell git config remote.origin.url)
	@echo
	@echo SHOUDL remote.upstream.url=git@github.com-gerardwebb:amplify-edge/amp-shared
	@echo remote.upstream.url: $(shell git config remote.upstream.url)
	@echo
	@echo
	@echo
	@echo

gitr-print-upstream:
	

## Prints the git setting
gitr-print: gitr-print-raw
	@echo
	@echo -- GITR Origin fork  --
	@echo GITR_ORG_ORIGIN: 				$(GITR_ORG_ORIGIN)
	@echo GITR_SERVER: 					$(GITR_SERVER)
	@echo GITR_USER: 					$(GITR_USER)
	@echo GITR_REPO_NAME: 				$(GITR_REPO_NAME)
	@echo
	@echo -- GITR Upstream --
	@echo GITR_ORG_UPSTREAM: 			$(GITR_ORG_UPSTREAM)
	@echo GITR_REPO_UPSTREAM_URL: 	$(GITR_REPO_UPSTREAM_URL)
	@echo
	@echo ---
	@echo GITR_REPO_ORIGIN_URL: 		$(GITR_REPO_ORIGIN_URL)
	@echo GITR_REPO_ORIGIN_FSPATH: 		$(GITR_REPO_ORIGIN_FSPATH)
	@echo
	@echo ---
	@echo GITR_VERSION: 				$(GITR_VERSION)
	@echo GITR_LAST_TAG:				$(GITR_LAST_TAG)
	@echo




### GIT-FORK

#See: https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/syncing-a-fork

gitr-fork-all: gitr-status gitr-fork-catchup gitr-status gitr-fork-commit gitr-fork-push gitr-fork-open

gitr-status:
	# Test
	git status


### FORK

## Prints the exact Git Clone command you need :)
# You can call this from your Org folder:
# make -f shared/boilerplate/gitr.mk gitr-fork-clone-template
gitr-fork-clone-template:
	@echo
	@echo
	@echo Template is:
	@echo EX "git clone git@github.com-ME-amplify-cms:ME-amplify-cms/REPO_NAME"
	@echo
	@echo So if your fork is: 
	@echo github.com/james-amplify-cms/dev
	@echo
	@echo You use: 
	@echo EX: "git clone git@github.com-james-amplify-cms:james-amplify-cms/dev"
	@echo

## Sets up the git fork locally.
gitr-fork-setup-old:
	# Pre: you git forked ( via web) and git cloned (via ssh)
	# add upstream repo

	#git remote add upstream git://$(GITR_SERVER)/$(GITR_ORG_UPSTREAM)/$(GITR_REPO_NAME).git

## Sets up the git fork locally.
gitr-fork-setup:
	# Pre: you git forked ( via web) and git cloned (via ssh)
	# Sets up git config upstreak to point to the upstream origin
	@echo
	@echo EX git remote add upstream git@github.com-amplify-cms:amplify-cms/dev
	@echo
	@echo EX git remote add upstream git@$(GITR_SERVER)-$(GITR_USER):$(GITR_ORG_UPSTREAM)/$(GITR_REPO_NAME)
	@echo
	# WORKS
	git remote add upstream git@$(GITR_SERVER)-$(GITR_USER):$(GITR_ORG_UPSTREAM)/$(GITR_REPO_NAME)
	@echo

GITR_SHARED_URL=https://go.amplifyedge.org/sys-v2
gitr-fork-submod-setup:
	# THis is to get shared into a repo as a sumodule. SO that its easy to do CI
	# DONT do this in SHared !! Will be crazy recursive
	git submodule add -b master $(GITR_SHARED_URL)
	git submodule init
gitr-fork-submod-update:
	git submodule update --remote
gitr-fork-submod-delete:
	# Steps: https://git.wiki.kernel.org/index.php/GitSubmoduleTutorial#Removal
	# Delete the relevant line from the .gitmodules file.
	# Delete the relevant section from .git/config.
	# Run git rm --cached path_to_submodule (no trailing slash).
	# Commit the superproject.
	# Delete the now untracked submodule files.

	git submodule deinit -f — mymodule
	rm -rf .git/modules/mymodule
	git rm -f mymodule


## Sync upstream with your fork. Use this to make a PR.
gitr-fork-catchup:
	
	# This fetches the branches and their respective commits from the upstream repository.
	@echo
	git fetch upstream
	@echo

	# This brings your fork's master branch into sync with the upstream repository, without losing your local changes.
	@echo
	git merge upstream/$(GITR_BRANCH_NAME)
	@echo

## Commit the changes to the repo
gitr-fork-commit:
	@echo GITR_COMMIT_MESSAGE: $(GITR_COMMIT_MESSAGE)
	git add --all
	git commit -m '$(GITR_COMMIT_MESSAGE)'

## Push the repo to orgin
gitr-fork-push:
	git push origin $(GITR_BRANCH_NAME)

## Opens the forked git server.
gitr-fork-open:
	open $(GITR_REPO_ORIGIN_URL).git

## Submits the PR you pushed
gitr-fork-pr-submit:
	## TODO. Alex gave me the commands.
	open $(GITR_REPO_ORIGIN_URL).git

### UPSTREAM

## Opens the upstream git web
gitr-upstream-open:
	open https://$(GITR_SERVER)/$(GITR_ORG_UPSTREAM)/$(GITR_REPO_NAME).git 

gitr-upstream-merge:
	# Use github cli




## GIT-TAG

## Create a tag.
gitr-tag-create:
	# this will create a local tag on your current branch and push it to Github.

	git tag $(GIT_TAG_NAME)

	# push it up
	git push origin --tags

## Deletes a tag.
gitr-tag-delete:
	# this will delete a local tag and push that to Github

	git push --delete origin $(GIT_TAG_NAME)
	git tag -d $(GIT_TAG_NAME)

## GIT-RELEASE

## Stage a release (usage: make release-tag VERSION={VERSION_TAG})
gitr-release-tag:
	@echo Tagging release with version "${VERSION}"
	@git tag -a ${VERSION} -m "chore: release version '${VERSION}'"
	@echo Generating changelog
	@git-chglog -o CHANGELOG.md
	@git add CHANGELOG.md
	@git commit -m "chore: update changelog for version '${VERSION}'"

## Push a release (warning: ensure the release was staged first)
gitr-release-push: 
	@echo Publishing release
	@git push --follow-tags