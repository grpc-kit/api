.PHONY: proto

proto:
	protoc -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/gogo/googleapis/ \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=:${GOPATH}/src/ \
		--go-http_out=:${GOPATH}/src/ \
		--go-grpc_out=:${GOPATH}/src/ \
		proto/v1/*.proto
