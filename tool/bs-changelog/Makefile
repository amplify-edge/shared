# OLD make file showing usage of the gen change log code.
# Needs work to inetgrate into hugo.
# We will publish to gcs or our own web site and NOT github, so using hugo for change log is needed.

## site-common
site-common: clean 
	if [ ! -d themes/hugo-theme-soloio ]; then git clone https://github.com/solo-io/hugo-theme-soloio themes/hugo-theme-soloio; fi
	# style updates for putting docs in the gloo repo, see details here https://github.com/solo-io/hugo-theme-soloio/commit/e0c50784a92fb7f61c635ff9a6e3a010f636f550
	git -C themes/hugo-theme-soloio checkout a9c18a63d56bea026a9e241ce0078caf56eabbc5
	GO111MODULE=on go run cmd/generate_changelog_doc.go gen-changelog-md gloo > content/static/content/gloo-changelog.docgen
	GO111MODULE=on go run cmd/generate_changelog_doc.go gen-changelog-md glooe > content/static/content/glooe-changelog.docgen

## site-test
site-test: site-common 
	GO111MODULE=on go run cmd/generate_changelog_doc.go gen-version-scope-data --no-scope
	hugo --config docs.toml
	# ensure that valid json is generated. Common cause: using yaml ">" multiline string symbols in Hugo's toml headers
	cat site/index.json | jq "." > /dev/null

## site-release
site-release: site-common 
	GO111MODULE=on go run cmd/generate_changelog_doc.go gen-version-scope-data --product gloo --version $(VERSION) --call-latest
	hugo --config docs.toml
	# ensure that valid json is generated. Common cause: using yaml ">" multiline string symbols in Hugo's toml headers
	# (if it passes here, it will pass on the subsequent generation so no need to check both hugo calls)
	cat site/index.json | jq "." > /dev/null
	mv site site-latest
	GO111MODULE=on go run cmd/generate_changelog_doc.go gen-version-scope-data --product gloo --version $(VERSION)
	hugo --config docs.toml
	mv site site-versioned

## site-serve
site-serve: site-test 
	hugo --config docs.toml --themesDir themes server -D

## clean
clean: 
	rm -fr ./site ./resources ./site-latest ./site-versioned ./data/Solo.yaml