# dev

So the SDK will be comiled and released in github actions.
SO then other repos can at the strat just grab it from the releases.
Then they can call it from a simple makefile or golang tasks (https://github.com/go-task/task)
	- with go tasks we remove the make file issues wth windows.

DB
- This has exactly what we want sys-core to have: 
	- https://github.com/cockroachdb/cockroach-go/tree/master/testserver
	- https://github.com/cockroachdb/cockroach/blob/master/pkg/server/testserver.go
	- OR
	- https://github.com/cockroachdb/cockroach/blob/master/pkg/cli/start.go
- this has a nice Text Fixtures for ti: https://github.com/go-testfixtures/testfixtures

Alex has the go-Generate working off the make files.
- SO lets keep everything that way for now.
- Can use go-getter: https://github.com/phogolabs/cloud/blob/master/pubsub/spec.go
	- If the SDK includes go-getter it can be used to do this.
		- e.g: https://github.com/nagu2k15/Minikube1/blob/master/pkg/util/progressbar.go

ci
- The way we setup the repo to then call make is broken.
- I saw a smarter way that works and so just need to fix it

gitr
- make it able to do it on all repos in one hit
- make it send a Telegram message. would be good NOT to have to rely on github.
	- so then when a dev does a gitr operation we all see it.

dep
- need the dep stuff to use golang for all OS's
- this will get CI working

flutter
- package refs should use git or file references ?
- file references are fine because then local and CI just work because we assume that the other git repos are checked out by the CI script
- And the SDK, will do that checking out for you.
	- We will have a special sdk.yaml in which you declate what other repos you need in order to build that are part of GCN Or other Module developers.
	- Or we could use https://github.com/jsonnet-bundler/jsonnet-bundler
		- this woudl solve the golang generate problem
		- it would also solve the K8 / k3d problem.

generator

- we need alex's go_geneate to call our own Gen golang binary
- then we can start to get the generator working off golang and not make files
- first will be the embed
- then the replacement of gen things in the makefile to use golang.

serving flutter
- this works well: https://github.com/shurcooL/vfsgen
	- note: https://github.com/shurcooL/vfsgen/issues/75


enception handling
https://github.com/cockroachdb/sentry-go/blob/master/example/basic/main.go
- running a sentry in the cloud somewhere is easy.
	- https://hub.docker.com/_/sentry
