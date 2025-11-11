package gateway

import (
	"context"
	"crypto/tls"

	gwHello "github.com/achtarudin/grpc-sample/protogen/hello/v1"
	gwResiliency "github.com/achtarudin/grpc-sample/protogen/resiliency/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterHandlerFromEndpoint(ctx context.Context, config *GatewayConfig) error {
	var opts []grpc.DialOption

	creds := insecure.NewCredentials()

	if config.GrpcTLS {
		creds = credentials.NewTLS(&tls.Config{})
	}

	opts = append(opts, grpc.WithTransportCredentials(creds))

	// Register gRPC Services Here

	// Resiliency Service
	if err := gwResiliency.RegisterResiliencyServiceHandlerFromEndpoint(ctx, config.ServeMux, config.GrpcRemoteServer, opts); err != nil {
		return err
	}

	// Hello Service
	if err := gwHello.RegisterHelloServiceHandlerFromEndpoint(ctx, config.ServeMux, config.GrpcRemoteServer, opts); err != nil {
		return err
	}

	return nil
}
