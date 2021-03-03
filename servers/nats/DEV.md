# nats

Nats Jetstream provides HA and non HA replication.

Uses for NATS
- database chaneg feeds. A DB or File system produces change feeds, allowing data synchronisation in general.
- workflows / activities to allow one thing to notify anyother thing. Essentially a message bus.

Embed
- Can NATS embeded talk to NATS Jetstream

Non Embed
- Nats Jetstream

Encoding
- Protocol buffers is preferred, and ehcne why we made our GRPC API to have Service and Models seperated, thus allowing the same APi to run over GRPC or NATS.

