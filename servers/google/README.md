# google

For global we use google Storgae and Google PUB Sub, so we dont have to run our own fully HA.

TIDB we have to run, unless we start to use Bigquery.

Its a weird mixture of dev, build and runtime tools that need it.

This code provides a unified way to access any google bucket.

It expects the caller to provide all auth, path and any other context


## LIBS

Manually using the go-cloud stuff.
Example https://github.com/google/go-cloud/tree/master/samples/order, shows Google storage and Google Pub Sub being used together which is nice because when a backup happens a pubsub event will fire to tell the server process that a file changed. This helps avoid race conditions because one process when finished fires an event.

https://github.com/creachadair/badgerstore and https://github.com/creachadair/gcsstore both implement the same interface from https://github.com/creachadair/ffs !
gcsstore uses https://github.com/google/go-cloud
Might be some opportunities here for backup and restore in general but also for other things.