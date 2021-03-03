# badger

genji provides a basic SQL APi

- needs migrations.
- needs changed feeds for Materialised views and for HA
- needs HA
 - can be done with NATS Jestream
 - can be done with  github.com/BBVA/raft-badger


**rony**
- bad boy that uses only badger and raft to DB and sync. No nats basically.
- full code gen from grpc model and server to DB and sync. Its impressivly concise.