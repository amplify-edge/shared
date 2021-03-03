RONY_BIN=$(PWD)/_bin/rony
RONY_GEN_BIN=$(PWD)/_bin/protoc-gen-gorony

DCLOCK_BIN=$(PWD)/_bin/dclock
DCLOCK_DATA=$(PWD)/_data/dclock



print:
	@echo
	@echo RONY_BIN: $(RONY_BIN)
	@echo RONY_GEN_BIN: $(RONY_GEN_BIN)

	
	@echo
	@echo DCLOCK_BIN: $(DCLOCK_BIN)
	@echo DCLOCK_DATA: $(DCLOCK_DATA)
	@echo

dep:
	git clone --branch v0.5.29 https://github.com/ronaksoft/rony
	git clone https://github.com/ronaksoft/dclock
dep-delete:
	rm -rf rony
	rm -rf dclock

dep-goget:
	go get -u github.com/ronaksoft/rony/cmd/rony
	go get -u github.com/ronaksoft/rony/cmd/protoc-gen-gorony
dep-goget-delete:
	rm /Users/apple/workspace/go/bin/rony
	rm /Users/apple/workspace/go/bin/protoc-gen-gorony
	

build:
	cd rony/cmd/protoc-gen-gorony && go build -o $(RONY_GEN_BIN) .
	#cd rony/cmd/rony && go build -o $(RONY_BIN) .
	
	#cd dclock/cmd/cli-dclock && go build -o $(DCLOCK_BIN) .

sample:
	# oopsy fails. I thinks its because of him usng go 1.5...
	mkdir -p example
	cd ./example && rony project create --project.name github.com/ronaksoft/example --project.name example-project
	
data-clean:
	rm -rf $(DCLOCK_DATA)

run-client:
	$(DCLOCK_BIN) client \
		--clientID client01 \
		--hostPort 127.0.0.1:81

run-server:
	# one
	#$(DCLOCK_BIN) server

	$(DCLOCK_BIN) server \
		--bootstrap false \
		--dataPath $(DCLOCK_DATA) \
		--gatewayAdvertiseUrl 127.0.0.1:91 \
		--gatewayListen 0.0.0.0:80 \
		--gossipPort 7081 \
		--idleTime 1m0s \
		--join node01 \
		--raftPort 7080 \
		--replicaSet 1 \
		--serverID node01 \
		--tunnelAdvertiseUrl 127.0.0.1:81\
		--tunnelListen 0.0.0.0:81
		
run-server-nodes:
	# many

	# tunnel: 192.168.2.100:81
	# server: 192.168.2.100:80