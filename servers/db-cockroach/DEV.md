# CR

Can we swap to CR

encryption at rest
- NOPE: Its only for Enterprise: https://www.cockroachlabs.com/docs/stable/encryption.html#encryption-at-rest-enterprise

Backup and Restore
- full backups are for Core:  https://www.cockroachlabs.com/docs/v20.2/take-full-and-incremental-backups#full-backups
	- restore is also for Core at least.
	- can backup single or cluster
	- can restore singel or cluster.
- incremental are onyl for Ent: https://www.cockroachlabs.com/docs/v20.2/take-full-and-incremental-backups#incremental-backups
	- so once you reach a large size your screwed.
- NOT encrypted for Core:
- Scheduling is possible with Core: https://www.cockroachlabs.com/docs/v20.2/manage-a-backup-schedule

We need Materilialised views:
https://www.cockroachlabs.com/docs/v21.1/refresh.html#required-privileges
- when source table change, you must call REFRESH on the Sink table !
- how expensive is it ????

We need change feed from the Materilialised views:
https://www.cockroachlabs.com/docs/v20.2/stream-data-out-of-cockroachdb-using-changefeeds.html
Basic: https://www.cockroachlabs.com/docs/v20.2/stream-data-out-of-cockroachdb-using-changefeeds.html#create-a-core-changefeed
- a golang long running process connects as root
	- cockroach sql --url="postgresql://root@127.0.0.1:26257?sslmode=disable" --format=csv
- stopping must be explicit.


