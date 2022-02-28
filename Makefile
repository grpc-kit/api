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

proto-gogo:
	protoc -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/gogo/googleapis/ \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
        --gogo_out=plugins=grpc,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types:\
${GOPATH}/src/ \
		proto/v1/*.proto
