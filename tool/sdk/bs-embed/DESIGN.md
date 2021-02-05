 # DESIGN

 This is the chronological order that embed will run


// 0. Dev time
// use the "bs-migrate" tool to scaffold and write your migrations.

// 1. Gen Time
// Just call "go generate" from Main.
// In each Module, generate.go exists with "go:generate bs-gen".
// Will the golang compiler do in bottom up order: sys-core, then sys-*, then mod-*, your main ?

// 2. Embed time ( maybe can be done at 1. Gen time also)
// Just call "go generate" from Main.
// In each Module, generate.go exists with "go:generate bs-gen".
// Each Module is creating its bindata binary file called "bindata-go", using the go-bindata tool
// Will the golang compiler do in bottom up order: sys-core, then sys-*, then mod-*, your main ?

// 3. Compile time
// Main now does a standard compile, based on its standard import paths.
// All the sub modules bindata is naturally included in the compile.

// 4. Config time
// Main needs to unpack some stuff just before its own build.
// calls a function in sys-core, with a string for the asset folder it wants.
// - each modules base config
// - flutter lang json files
// - flutter base config (todo)

// 5. Run Time
// Now when sys-core needs the packed assets it can get them
// - auto migrations Or manual migrations ( via CLI or Flutter GUI )
// 	- Do from bottom up ( sys-core migrations first, etc etc)


---

https://github.com/EmbarkStudios/wg-ui/blob/master/Makefile
assets:
	$(GOGET) github.com/go-bindata/go-bindata/...
	$(GOGET) github.com/elazarl/go-bindata-assetfs/...
	go-bindata-assetfs -prefix ui/dist ui/dist

---

https://docs.iris-go.com/iris/file-server/introduction
- shows easy way to use go-bindata

---

github.com/shurcooL/vfsgen
https://git.furworks.de/opensourcemirror/Gitea/commit/83b90e41999d30e4abb46f6bf0f1c3359cfd4d04
