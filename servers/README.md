# Test Harness

Tools to make it easy to boot up your environment from your laptop for local dev, but also your servers.

This makes Dev, CI, CD and OPS easier and repeatable.

All these tools will be refactored to be used by Booty.

TODO:

- everything here needs to be converted to golang so that booty can run it. Makefile for now just to work out how to do it.



## Main 

How does this relate to main binary ?

Our main binary is designed to run as a single binary with nothing else, and so allows a user to run everything on a desktop or linux server easily.

It really can be thought of as the Edge binary because it has a DB and GUI, and can be easily packaged and run by anyone. Its currently is designed to operate on it owns.
But later, it can start to act as a CDN in that read data can be cached in the DB, and mutations can be saves locally and forwarded to a Server in a Event Architecture.

Main needs the following stuff from the Server.
- HA or Scalability of that main binary (v2 roadmap) or many binaries (v3 roadmap)
- Ops tools.
- Telemetry.

The tools below give us these things we need.

## Autostart

Systemd is ok, but we can also use one that works on Desktops and Servers. 
- We need both !

Desktops: https://github.com/emersion/go-autostart

## Config

Because of the modular architecture we have many separate configs.

Alex has proposed a Config wrapper to "merge" the config all into one.
But we also need to be able to alter the config for users ( like a cli wiazrd or other ), and to be able to migrate configs just like we migrate a DB.

SO instead it woudl be better to not "Merge" the config but instead write a wrapper that can take a standard config, and manipulate it.
When this is teamed up with Booty, it will keep things modular and give us a good onboarding architecure.

## Secrets

We currently store them in systemd, which is not gong to stop a hacker that gets physical access.

Instead we need to store secrets in the TPM chip.
https://en.wikipedia.org/wiki/Trusted_Platform_Module


- Linux and windows
	- https://github.com/salrashid123/tpm2
	- build on google project.
- Linux, Windows and Mac
	- https://github.com/zalando/go-keyring
	- wrapper: https://github.com/martinohmann/keyring
- Cloud
	- Use go-cloud lib.

## Control Plane : Deployment And backup

We want to use google storage to hold config and bootstrap so its easy to redeploy.

Github Server can also be used, we need a Proper storage server anyway, so we think its better to store configs in git, push them to Google Storage and then sysnc Servers with Google Storage.
SO then your github releases can just push zip everything and push it to Google Storage.
Mini can replace google storgae if less HA is needed later.

DEV :We just use Yurt but extend it.
- Nomad runs starts it up
- https://www.nomadproject.io/docs/drivers/exec
- https://www.nomadproject.io/docs/drivers/raw_exec
- Yurt can be 

FLow:

Booty is copied to laptop and sever to start the flow...

- Booty
	- Put booty on server ( ssh in a pull it via Curl.)
		- Call Get binaries
			- main
			- caddy
			- minio
			- nats
		- Start it ALL up with a systemd.
			- booty can start the others and manage them.
		- config
			- booty.yaml has just a token plus the bit for:
			- minio needs 5 things.
			- nats needs 2 things.
	- CLI Connects with the token.
		- pass in the URL ( local or remote ).
		- pass in the token.
	- Stage 0: Run it all locally !!!
		- Get your configs all working, etc.
		- Provision to the local MINIO.
	- Stage 1 Prov: Pushes the everythig to the server as a zip.
		- binaries
		- systemd
		- config
			- the real security logons.
	- Stage 2: Same thing to google
		- Booty Server gets the event and ignore or ...
- Then en day the server falls over..
	- SO do stage 2 in reverse.
	- then so Stage 1 again to the server.
- backups
	- Main is doing daily backups...
	- Booty is watching for them on the File system
	- Booty then backs it up to Google OR our own Minio :)


Mechanisms:
- no code gen. nothing fancy.
- cobra
- file movements: filehelper.go copied to booty.
- talk to google with gocloud.dev/blob/<driver>
- talk to nats and gogole pub sub with gocloud.dev/pubsub/<driver>


## Booty

How does tis relate to booty ?

Booty is what a Developer or User needs to deploy and run the system, and so booty will control the tools using shell commands.

It gives us the Ops ( Operations ) tooling that devs and users need to manage the full life cycle of Development, Deployment and Updates for all the things needed to run the system.

Booty-Dev is for developers

- imports booty cmd, so it gets that too.
- imports the code git, code gen tools from shared.

Booty is for users

- just has the stuff to deploy and admin the runtime binaries.


TODO:

- TIUP has nicer signal control in the terminal, and i think we can use that in booty to control all the binaries...
- Use an agnsotic Lib to talk to the other things in booty and our main binary.
	- https://github.com/google/go-cloud (https://gocloud.dev/)
		- Makes S3 agnostic ( gocloud.dev/blob )
		- Makes NATS agnostic.
		- Makes TIDB agnostic ( gocloud.dev/mysql )
		- Makes Hashicorp Vault agnostic ( gocloud.dev/secrets )
			- Not sure we need it.
			- On baremetal, where to store secrets ?. We need to use Zoolando golang code i think so it woks on all Desktops and Servers to store secrets in the native TPM chip.

- Yurt has a really nice way of pulling binaries. Use that in Booty and otehr tools.


## Hashicorp

Gives us a Deployment Orchestration basis for exe and docker.

Based on: https://github.com/ncabatoff/yurt. Yurt boots up the following binaries: nomad, consul, vault, prometheus, consul_exporter, node_exporter.

Users or we can use this to deploy to Server(s) and manage their lifecycle.

Bootstrapping itself:

- Booty embeds the yurt code to install the hashicorp bits.
- Booty can use a global s3 / pubsub to know all servers running via the meta data.
- Booty can manage upgrades via this meta data.

Deploying other binaries / dockers:

- There are a few options for how we can use this.
- Use the "Raw Exe" functionality of Nomad to deploy any binary to a Desktop or Server.
- Use the "docker" functionality of nomad to deploy dockers with docker swarm.
	- this might be a perfect use case for Module, because they are untrusted, and our main binary can use the modules over GRPC.

Caddy can be on the Server and be our main host.
Caddy will then just reverse proxy all docker deployed using https://github.com/lucaslorentz/caddy-docker-proxy/


## Kubernetes

We wont need k8 for a long time, and maybe never.

openyurt ( https://github.com/alibaba/openyurt ) hwoever is a good example of tooling to make it easy for devs and users to work with k8.

## victora metrics

Stack:

Caddy
- For basic security needs.


Main binary (org server)
- protocols
	- Standard prometheus
		- we use github.com/VictoriaMetrics/metrics
	- standard loggin lib: zap or logrus in json format ?
		- we use zap.
	- Exposes an endpoint for the agent to pull from.
		- /metrics
		- none.

VicMetAgent ( vmagent runs on our server)
- pulls from Main binary
- polling ( 10 secs )
- can pull from many servers....

VicMet Store ( run with Grafana )
- File system dir using custom format.

VicMetBackup ( run with Grafana )
- pulling from VicMet Store ( streaming ..)
- pushing to S3.
- Sec
	- One per Org. So we give it out to them and they put in their config.

S3 ( Google GSC)
- Bucket per Org.

Grafana ( Our server  )
- polling the agent store ( 10 secs )
- Grafana
	- Sec
		- One for org. So we give it out to them and they store wherever.
		- One for Biz Team who have get "all labels view". So we give it out to them and they store wherever.
	- Proxy Sec ( if it works and 5 days, so not doing. )
		- https://grafana.com/docs/grafana/latest/auth/auth-proxy/
		- SO we provide a Link that has the headers in it from our Sys-accounts...
- Caddy
	- https://amplify-cms.org/metrics
		- Sec - None.


Maybe later...
- https://github.com/lindenlab/caddy-s3-proxy
- https://github.com/trusch/caddy-extauth


## grafana

Gives us a Telemetry aspects we need.

TempoDB looks perfect for us because it uses a simple golang WAL store, and then sends it all into S3. It has not other moving parts except for a few binaries.
It is based on JSONET to make configuration easy, which we use for main architecture.


TODO:
- There is no booty for the grafana toolset, and so we have to make it.
- Get booting of grafana to be booty like. TIUP has code to do it that we can use.
- Grafana GUI is good enough for Ops users, and we will make one in Flutter charts for Biz users dashboards.


## tidb

Gives us a SQL DB that can scale out with the CDC aspects that our CRS architecture needs.

TODO:
- It boots and uses grafana, but we need to not use it but instead our own Grafana booty.

## minio

Gives is a Files store that cna scale out.

TODO: 
- Make file works, and has basic E2E make based tests.
- Get it integrated into NATS, so we get a CDC of Minio.
- 
## nats

Gives us a Message Queue and Event Bus.

TODO:
- THere s no booty for it, and so we need to make that.
- Integrate with TIDB and Minio so that those 2 stores have CDC.
