print:
	cockroach version

	grep 'node starting' node1/logs/cockroach.log -A 11
	grep 'node starting' node2/logs/cockroach.log -A 11
	grep 'node starting' node3/logs/cockroach.log -A 11

dep:
	brew install cockroachdb/tap/cockroach

db-run:
	cockroach start \
		--insecure \
		--store=node1 \
		--listen-addr=localhost:26257 \
		--http-addr=localhost:8080 \
		--join=localhost:26257,localhost:26258,localhost:26259 \
		--background
db-run-nodes:
	cockroach start \
		--insecure \
		--store=node2 \
		--listen-addr=localhost:26258 \
		--http-addr=localhost:8081 \
		--join=localhost:26257,localhost:26258,localhost:26259 \
		--background

	cockroach start \
		--insecure \
		--store=node3 \
		--listen-addr=localhost:26259 \
		--http-addr=localhost:8082 \
		--join=localhost:26257,localhost:26258,localhost:26259 \
		--background
db-init:
	# must run this after ndoes are up in order for Web GUI and SQL to work.
	cockroach init --insecure --host=localhost:26257

db-stop:
	#cockroach quit --insecure --host=localhost:26257
	cockroach node drain --insecure --host=localhost:26257
db-stop-nodes:
	cockroach node drain --insecure --host=localhost:26258
	cockroach node drain --insecure --host=localhost:26259

db-clean:
	rm -rf node1 
db-clean-nodes:
	rm -rf node2 node3

db-cli:
	cockroach sql --insecure --host=localhost:26257

db-web:
	open http://localhost:8080/
	
db-sample:
	# each Node uses about 100 MB and 2 % of CPU at idle. Its workable.
	cockroach workload init movr \
		'postgresql://root@localhost:26257?sslmode=disable'
db-sample-long:
	# each Node uses about 120 MB and 30 % of CPU while running. Also workable.
		--duration=5m \
		'postgresql://root@localhost:26257?sslmode=disable'