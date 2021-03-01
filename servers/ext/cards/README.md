# cards

Cards is a standard way to describe GUI Forms that are also able to pist data back.

There are also many other approaches too, like describing the GUI layout with JSON, but htey generally dont include a Postback.

We dont want to surface our Gsuite stuff via this, because those things are designed to be backed in to the compiled app and then called with params to contextualise the tenant, user, context aspects. One to one mapping to the golang behind it.

Cards allow the Server to drive the GUI
- Devs and Users can quickly make a GUI.
- no compile or deploy.
- useful for data driven systems like Admin and mod-disco possible.
- when hooked up to our server supporitng CQRS and Functions, it will give a good no code development on top of the modules and widgets we are buildng.

Web:https://github.com/Microsoft/AdaptiveCards/

Schema:https://adaptivecards.io/explorer/

vscode:

- https://marketplace.visualstudio.com/items?itemName=madewithcardsio.adaptivecardsstudiobeta
- https://marketplace.visualstudio.com/items?itemName=tomlm.vscode-adaptivecards



flutter Rendering engine

- https://pub.dev/packages/flutter_adaptive_cards
	- Works very well, and is simple to use.

golang:
- github.com/DanielTitkov/go-adaptive-cards v0.2.1