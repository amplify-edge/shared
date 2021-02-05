# dev

examples where minio and nats maatch Google storage and Google Pub Sub.

https://github.com/Gerifield/mini-asciinema-store
- only blob

https://github.com/lsmith130/deq
- only blob, but server and cli at least
- has a custom code generator based on proto too.
- store. uses badger: https://github.com/lsmith130/deq/blob/master/deqdb/store.go
	- has upgrader too: https://github.com/lsmith130/deq/blob/master/deqdb/internal/upgrade/upgrade.go