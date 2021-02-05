# grpc tools

# This puts everything onto the global path.
# Assumes you have Golang installed.
# Assumes you are on a MAC.


# Found from : https://github.com/lynxkite/lynxkite/blob/master/sphynx/proto_compile.sh

# Envoy
LIB_ENVOY=						github.com/tetratelabs/getenvoy
LIB_ENVOY_REPO_FSPATH=			$(GOPATH)/src/$(LIB_ENVOY)
LIB_ENVOY_GETENVOY_VERSION= 	v0.1.8
LIB_ENVOY_VERSION= 				standard:1.14.3

# grpcui is a golang Web GUI for GRPC
# https://github.com/fullstorydev/grpcui
LIB_GRPCUI_REPO=				github.com/fullstorydev/grpcui
LIB_GRPCUI_REPO_FSPATH=			$(GOPATH)/src/$(LIB_GRPCUI_REPO)

# GO GRPC is now here:https://github.com/grpc/grpc-go
LIB_GOGRPC_REPO=				github.com/grpc/grpc-go
LIB_GOGRPC_REPO_FSPATH=			$(GOPATH)/src/$(LIB_GOGRPC_REPO)
LIB_GOGRPC_REPO_VERSION= 		v1.32.0

# New
# https://github.com/protocolbuffers/protobuf-go
# which is https://godoc.org/google.golang.org/protobuf/cmd/protoc-gen-go
LIB_GO_REPO=					github.com/protocolbuffers/protobuf-go
LIB_GO_REPO_FSPATH=				$(GOPATH)/src/$(LIB_GO_REPO)
LIB_GO_REPO_VERSION= 			v1.25.0

# Old ( last update in 14 May, 2020, so very old )
# https://github.com/golang/protobuf/
LIB_GOOLD_REPO_REPO=			github.com/golang/protobuf
LIB_GOOLD_REPO_REPO_FSPATH=		$(GOPATH)/src/$(LIB_GOOLD_REPO_REPO)

# grpc-gateway & swagger
# https://github.com/grpc-ecosystem/grpc-gateway
# has github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
# has github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
LIB_GRPC_GATEWAY_REPO=			github.com/grpc-ecosystem/grpc-gateway
LIB_GRPC_GATEWAY_REPO_FSPATH=	$(GOPATH)/src/$(LIB_GRPC_GATEWAY_REPO)
LIB_GRPC_GATEWAY_REPO_VERSION=	v1.15.0

# Protoc
# All tags here: https://github.com/protocolbuffers/protobuf/tags
# For now just using Brew. Will upgrade to use a golang downloader soon.
LIB_PROTOC_VERSION=				3.13

## Print
grpc-print:
	@echo
	@echo -- Envoy --
	@echo LIB_ENVOY_REPO_FSPATH: 		$(LIB_ENVOY_REPO_FSPATH)
	@echo LIB_ENVOY_GETENVOY_VERSION: 	$(LIB_ENVOY_GETENVOY_VERSION)
	@echo LIB_ENVOY_VERSION: 			$(LIB_ENVOY_VERSION)

	@echo
	@echo -- GRPCUI GUI --
	@echo LIB_GRPCUI_REPO_FSPATH: 		$(LIB_GRPCUI_REPO_FSPATH)

	@echo
	@echo -- GO GRPC compiler --
	@echo LIB_GOGRPC_REPO_FSPATH: 		$(LIB_GOGRPC_REPO_FSPATH)
	@echo LIB_GOGRPC_REPO_VERSION: 		$(LIB_GOGRPC_REPO_VERSION)
	
	@echo
	@echo -- GO Protobuf compiler --
	@echo LIB_GO_REPO_FSPATH: 			$(LIB_GO_REPO_FSPATH)
	@echo LIB_GO_REPO_VERSION: 			$(LIB_GO_REPO_VERSION)
	
	@echo
	@echo -- OLD Golang Protobuf compiler --
	@echo LIB_GOOLD_REPO_REPO_FSPATH: 		$(LIB_GOOLD_REPO_REPO_FSPATH)

	@echo
	@echo -- GO GRPC Gateway compiler --
	@echo LIB_GRPC_GATEWAY_REPO_FSPATH: $(LIB_GRPC_GATEWAY_REPO_FSPATH)
	@echo LIB_GRPC_GATEWAY_REPO_VERSION: $(LIB_GRPC_GATEWAY_REPO_VERSION)
	
	@echo
	@echo -- Protoc compiler --
	@echo LIB_PROTOC_VERSION: 			$(LIB_PROTOC_VERSION)

	@echo
