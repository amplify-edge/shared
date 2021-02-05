# flu-proto

## Protoc which
flu-proto-which:
	@echo
	@echo -- flu-proto : which --
	flutter pub global list 
	@echo

## Protoc dep
flu-proto-dep:
	# http://google.github.io/proto-lens/installing-protoc.html
	@echo
	@echo -- flu-proto : dep --
	flutter pub global activate protoc_plugin

## Protoc-git-delete
flu-proto-delete:
	@echo -- flu-proto : delete --
	flutter pub global deactivate protoc_plugin

flu-proto-vscode:
	# NONE