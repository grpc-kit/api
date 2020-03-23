.PHONY: proto

proto:
	protoc -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/gogo/googleapis/ \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gogo_out=plugins=grpc:. proto/v1/*.proto
