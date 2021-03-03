# github.com/DanielTitkov/go-adaptive-cards

FLU_FSPATH=Flutter-AdaptiveCards/example

GO_BOT_FSPATH=teams-bot-example

print:
	@echo GO_BOT_FSPATH: $(GO_BOT_FSPATH)
	

### DEP

dep:
	# Flutter Code that parses Adaptive Card JSON, and displays the Flutter widget.
	git clone https://github.com/neohelden/Flutter-AdaptiveCards

	# Golang Server that uses Facebooks Ent driver and a SQL database.
	# Serves Adaptive Card and model the dialogue flows.
	# Designed to work with Microsoft Teams and Bot SDK https://github.com/microsoft/botframework-sdk, whch is analogues to our Orgs / Projects concept.
	# The golang implemenation is used internally is: github.com/infracloudio/msbotbuilder-go
	# Use Ent ( https://github.com/facebook/ent ) and Ent graphql (h ttps://github.com/facebookincubator/ent-contrib )
		# See: https://entgo.io/docs/graphql/

	git clone https://github.com/DanielTitkov/teams-bot-example

dep-delete:
	rm -rf Flutter-AdaptiveCards
	rm -rf teams-bot-example


## RUN
flu-print:
	flutter devices
flu-run-web:
	# works . Video also...
	cd $(FLU_FSPATH) && flutter run -d chrome
flu-run-web-release:
	# works . Video also...
	cd $(FLU_FSPATH) && flutter run -d chrome --release

flu-run-desktop:
	# works
	# need to flush pof file sometimes
	cd $(FLU_FSPATH) && rm macos/Podfile
	cd $(FLU_FSPATH) && flutter run -d macOS
flu-run-desktop-release:
	# needed because so slow with debug. was fast a month 
	cd $(FLU_FSPATH) && flutter run -d macOS --release


flu-run-mobile-init:
	# edit profile in xcode
	open $(FLU_FSPATH)/ios/Runner.xcodeproj
flu-run-mobile:
	# works
	cd $(FLU_FSPATH) && flutter run -d bdf90dc799709a013a25d0fc2df80e441df026f3
flu-run-mobile-release:
	# works
	cd $(FLU_FSPATH) && flutter run -d bdf90dc799709a013a25d0fc2df80e441df026f3 --release
flu-run-mobile-build:
	# change to ios 11 in xcode for runner and pods
	cd $(FLU_FSPATH)/ios && pod install
	#cd $(FLU_SAMPLE_FSPATH) && flutter build ios

	cd $(FLU_FSPATH) && flutter build ios --release --no-codesign



go-db:
	# get ent for db
	go get entgo.io/ent/cmd/ent

	# get tidb, so we dont need docker.
	cd ./../tidb && $(MAKE) run-play
go-run:
	cp dev.yml $(GO_BOT_FSPATH)/configs
	cd $(GO_BOT_FSPATH) && $(MAKE) run