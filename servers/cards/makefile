# github.com/DanielTitkov/go-adaptive-cards

FLU_FSPATH=Flutter-AdaptiveCards/example

GO_BOT_FSPATH=teams-bot-example

print:
	

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
flu-run-desktop:
	# works
	cd $(FLU_FSPATH) && flutter run -d macOS
flu-run-mobile-init:
	# edit profile in xcode
	open $(FLU_FSPATH)/ios/Runner.xcodeproj
flu-run-mobile:
	# works
	cd $(FLU_FSPATH) && flutter run -d bdf90dc799709a013a25d0fc2df80e441df026f3

go-run:
	cd $(GO_BOT_FSPATH) && go run .