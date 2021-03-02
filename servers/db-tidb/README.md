# tiup

https://docs.pingcap.com/tidb/stable/maintain-tidb-using-tiup
https://github.com/pingcap/tiup

https://github.com/go-kratos


The two things above give use lots of what we need for the whole architecture

- Many binaries but its ok, because it tiup has good tooling to boot it up and stop it.
	- Will just mean we need good desktop packaging for Usrs because we must be able to submit it to app store.
	- Will mean signing many binaries.
	- Not sure Grafana, which is unpacked at runtime will be OK in terms of signing on a mac... 
		- needs investigation.

- Patching
	- https://docs.pingcap.com/tidb/stable/maintain-tidb-using-tiup#replace-with-a-hotfix-package
		- Need for dev time too so wokring is fats

- CQRS and CDC basis
	- We know we need this and TIDB is pretty muc the best. Better than cockroachdb in many ways.
	- https://github.com/pingcap/ticdc
	- Need to add NATS and LiftBridge and we have a CQRS basis
	- Need to chekc Materialised Views in TIDB and SQL Parser.
		- Cant do it: https://docs.pingcap.com/tidb/stable/views#limitations
		- SO just use Tables, and implement our a Query parser to check if a Mutation from the CDC affects a SQL Query or not.
			- If it does then we just need to emit the query result set as a CUD Event to the Client.
	- Can we do encryption at rest ?
		- Check..

- Its a Cluster discovery and ops and install tool all in one.
	- Exactly what we were needing.

- Its got Prom and Graphana
	- And the crazy config it needs too.
	- We can extend to have victoria metrics and Loki.

- Files DB
	- There is no reason we cant use TIDB to store the chunked versions of the Files
	- And use the TIDB cluster aspects to do what SeeweedFS does and have the Directory of the files held in TIDB also.
	- In this way the location of the files and the file chunks are decoupled. Just like Seeweedfas does.
		- see: https://github.com/geohot/minikeyvalue

- But how easy to extend TIUP ?
	- Need our Sys and Modules 
	- Need Victoria Metrics
	- JSONNET / Config

- TIDB and CI / CD
	- Need a Gitops, Flagger setup.

- go-kratos
	- It also use TIDB 
	- It has nicely packaged the code gen for GRPC, GRPC-Web, and a few other things, and so we can extend from it
		- Will be pat of Booty repo.
		- Can wrap with mage, so that CI and CD and use it.
	- Has a project creator
		- https://github.com/go-kratos/service-layout#create-a-project
		- Needed for our Modules.

- gui
	- https://pub.dev/packages/pluto_grid
	- definitly the right fit as its excel with type bindings, and so can link backwared to the data model to edit data in a MView but also the sources in the flow alos.

- VPP ( virtual POwer plant)
	- Its exactly the same as a CQRS compute system, and what we will do with EU.
	- IEC 61850-7-420
	- https://www.vde-verlag.de/iec-normen/preview-pdf/info_iec61850-7-420%7Bed1.0%7Db.pdf
	- https://de.wikipedia.org/wiki/Virtuelles_Kraftwerk
	
