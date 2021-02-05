# booty

This is for Users to help them deploy, configure, backup and update servers.

It works off a global File Server and Message Bus allowing you to manage many servers.

It can be run on your laptop and / the Server, and they will work together to allow managing servers remotely.

## Binaries deployed

Booty runs as a daemon and manages the following bimaries and their startup.

Proxy Server (Caddy)

- Main Forwarding Proxy exposing the virtual hosts.

Main

- The main binary.

File System (Minio)

- The File Server where all configurations, backups, data bootstraps are stored

Messaging System (Nats)

- The messaging system that allows Booty to get events on file system changes so that it can react to changes

## How to use

The CLI gives the following functions.

Connect - URL, Token.

Info - Gives the current status, Lists the config files, etc. Configuration default are designed for the CLI and Server to be on the same laptop.

Bootstrap - Installs the needed binaries.

Up - will bring up the binaries.

etc.... See the Servers readme for the rest.
