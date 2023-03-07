.PHONY: proto

proto:
	protoc -I. \
		-I/usr/local/include/ \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I${GOPATH}/src/github.com/googleapis/googleapis/ \
		--go_out=. \
		--go-grpc_out=require_unimplemented_servers=false:. \
		known/example/v1/*.proto
	protoc -I. \
		-I/usr/local/include/ \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I${GOPATH}/src/github.com/googleapis/googleapis/ \
		--go_out=. \
		--go-grpc_out=require_unimplemented_servers=false:. \
		known/status/v1/*.proto
