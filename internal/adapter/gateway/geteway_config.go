package gateway

import "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

type GatewayConfig struct {
	GrpcRemoteServer string
	GrpcTLS          bool
	ServeMux         *runtime.ServeMux
}
