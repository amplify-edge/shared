# Sdk

We have arrived at essentially a SDK.

We can now bundle many of the different bs-* tools in one if we want

However because the sdk will need to call the running sys (2 binary deploy) OR main (single binary deploy) for migrations we need to be careful.
- for migrations the SDK should use the GRPC API, because its independent of the binary name
- But it must use the CLI to get there, and so we need to depend apon the sys GRPC code 


## Design

 This is the chronological order that the build will run in per modules


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


## Dev Time

Migrate
- Need to talk to the running sys system in order to get the sql dump.
- So we use the GRPC to CLI, and add the extra bits we need for the local stuff
- Local cli stuff:
	- https://github.com/gobuffalo/soda
	- https://github.com/gobuffalo/fizz
		- rip out what we want for file system skeleton and CLI
		- See the makefile example in the data folder. Its nice !

In order to put their migrations into the Developers local DB, they will need to use the sys-core CLI.
With test data.

The migrations are run against the DB, and NOT against the golang struct.
- This is because a V2 migration in SQL, will break if it relies on a V3 golang model struct.

Use a UTC datetime stamp ( for ordering )
- The tool will do it for you. Part of the shared sys-main cli or bs-data .
- We will use a timestamped folder to make it clean for migration.


Test Data

- generated against the golang model.
- embedded, so can be used to ensure a migration works and to hav the system filled with data.
- So you can write golang tests off this with no mocking
	- The tests run on the Test DB.

Real Data

- Hand generated generally
- Can be included in the embed and unpacked at Main.



## Gen Time

We need a golang binary that does what the make files did, and that the go:generate can call.

- If sys-core cant compile then you cant do generation, and so cant work on sys-accounts, etc
- Hence why we have shared !!
- So then use shared/bs-gen, and it has no dependency on sys.
- Boilerplate .mk files can stay there as a backup for now

CLI
- A Developer will need to run certain things partially, and not as part of a full gen.
- So the binary needs a CLI
- It always assumes the Directory its run from is the Dir to do the work on. Keeps things simple.
- SO the "generate" part of the CLI, which knows the directors structure we use will just call those sub CLI commands.

## Compile Time

Now its all easy. Its just a go build from Main.

## Config Time

Now you run the main binary, and can.

## Run Time

Sys-core now can do migrations, put the test data and real data into the DB.

Discovery

- It will reflect off embedded and the File system (where the binary is) to tell a dev or Ops what they have available.
	- A user can upload at runtime a new Real Data JSON also
- It can do it via the cli or GUI.
- It can tell you what is available
- It will use a Migrations Table to audit log all migrations and test data things that happened.
	- So it can return this to the usr also.
- It will enforce AuthZ

Running

- When it runs a Migration, it MUST order then globally.
	- So it must parse all the migrations available and then order then globally.
		- This is because a dev could have modified a Mod and Sys Module at different times and once is dependent on the other.
	- When the Server boots, if in Auto migration mode it must run them from where the last one is all the way to Present.

When it runs a Data load ( Test or Real )
	- It just loads it. The Test Data in not like migrations.
	- The Test data is always assumed to be designed to work with the Latest Migration.



