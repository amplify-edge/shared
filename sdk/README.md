# sdk

SDK / CI / CD / Update

Makefile are not cutting it anymore..Keep them while we make the transition

Unified. Use Yurt and extend the CMD
Hashicorp tools based
Can run our CI and CD locally or in Prod. Works like a champ.
Can run the CI and CD locally too
Gitea
Need a SQL DB
Can include in dep.
DNS (Cloudflare). Register automatically on boot
SEC.
Configs Firefox, 
Config Client OS
Config Server
Deploy
harcode to Hertzner for now. Add more later
SDK
Gen and embed and migration.


https://github.com/ncabatoff/yurt
- for nomad, consul, vault
- makes life easy.

https://github.com/alibaba/openyurt. 
- for k8


DB code gen and migration
https://github.com/kyleconroy/sqlc

lap
https://github.com/oragono/oragono-ldap

esni for client and server
- Compilng go for modules such that FFI works with flutter
https://github.com/iyouport-org/relaybaton/blob/master/Makefile