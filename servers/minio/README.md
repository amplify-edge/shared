# test harness

TempoDB can use any s3 backing store.

And we need an s3 backing store for:
- TempoDB ( if we run it ourselves )
- Deploymnet of binaries and channels
- Deployment of remote config and bootstraps.
- Maybe a decent File store for global stuff.

So if we use the google.cloud golang libs in our main code we can use either Google or our own to run things.

- Google Storge ( s3) with Google PubSub
- Minio with NATS Pubsub

These can be run under Caddy.

## Deployment in general

We have a single binary that runs on anything, so that gets use a Desktop
We need a HA setup ( like 3 servers) on hertzner or raspi at home.

So why not just use booty to download them, and get the parts out of tiup just for starting them up, upgrading them and shitting them down cleanly on any OS ?

We can also then use Yurt which runs Nomad and Consul, to manage all these binaries on many global servers if we want.
Or users to manage their own 3 servers.
https://github.com/ncabatoff/yurt

I got it wrapped in github.com/amplify-cms/shared/sdk ages ago..



