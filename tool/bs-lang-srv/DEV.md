# needs:


- relay
  - that is easy to run with caddy and disposable.
  - https://github.com/brave/go-translate

- api needs to be stable.
  - in app, uses GRPC
  - out of app, uses GRPC.

- translation memory so that less hits occur and we only hit the translation service if a cache miss occurs
   - make sure we can delete bad cache rows.
  
- providence 
  - make sure our API stores and returns the source of the translation, so bad translations can be overrideen and so the translation memory gets smarter over time. This is why gsheets are nice as there is a GUI, and a full providence source.

- tiered store for the translation memory and cache misses.
   - local hits genji ( memory or badger based), which then hits global ( google sheets ), whch then hits relay and original service.
   - so Dev / CI / App can all use it


