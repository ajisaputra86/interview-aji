//go:generate protoc -I protobuf --go_out=./model --go_opt=paths=source_relative --go-grpc_out=./model --go-grpc_opt=paths=source_relative protobuf/model.proto
//go:generate protoc -I protobuf --go_out=./internal/server --go_opt=paths=source_relative --go-grpc_out=./internal/server --go-grpc_opt=paths=source_relative protobuf/test.proto
//go:generate protoc -I protobuf --grpc-gateway_out ./internal/server --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt grpc_api_configuration=protobuf/test_grpc_gateway.yaml protobuf/test.proto
package main

import (
	"github.com/fajrirahmat/interview-aji/internal/server"
)

func main() {
	srv, _ := server.New()
	srv.Run()
}
