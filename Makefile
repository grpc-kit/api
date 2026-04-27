.PHONY: proto

proto:
	protoc -I. \
		-I/usr/local/include/ \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I${GOPATH}/src/github.com/googleapis/googleapis/ \
		--go_out=${GOPATH}/src \
		--go-grpc_out=require_unimplemented_servers=false:. \
		known/example/v1/*.proto

	protoc -I. \
		-I/usr/local/include/ \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I${GOPATH}/src/github.com/googleapis/googleapis/ \
		--go_out=${GOPATH}/src \
		--go-grpc_out=require_unimplemented_servers=false:. \
		known/status/v1/*.proto

	protoc -I. \
		-I/usr/local/include/ \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I${GOPATH}/src/github.com/googleapis/googleapis/ \
		--go_out ${GOPATH}/src \
		--go-grpc_opt require_unimplemented_servers=false \
		--go-grpc_out ${GOPATH}/src \
		./known/admin/v1/*.proto

	protoc -I. \
		-I/usr/local/include/ \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I${GOPATH}/src/github.com/googleapis/googleapis/ \
		--grpc-gateway_opt grpc_api_configuration=./known/admin/v1/admin.gateway.yaml \
		--grpc-gateway_out ${GOPATH}/src \
		--openapiv2_opt disable_default_errors=true \
		--openapiv2_opt disable_service_tags=true \
		--openapiv2_opt grpc_api_configuration=./known/admin/v1/admin.gateway.yaml \
		--openapiv2_opt openapi_configuration=./known/admin/v1/admin.openapiv2.yaml \
		--openapiv2_out=json_names_for_fields=false:./ \
		./known/admin/v1/admin.proto

	cp ./known/admin/v1/admin.gateway.yaml ${GOPATH}/src/github.com/grpc-kit/pkg/api/known/admin/v1/admin.gateway.yaml
