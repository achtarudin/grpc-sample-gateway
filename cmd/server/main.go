package main

import (
	"context"
	"fmt"
	"grpc-sample-gateway/internal/adapter/gateway"
	"grpc-sample-gateway/internal/adapter/http/handler"
	"grpc-sample-gateway/internal/adapter/logging"
	"grpc-sample-gateway/internal/helper"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(&logging.Format{})

	// Load .env file if exists
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error loading .env file, proceeding with environment variables.")
		return
	}

	// Get environment variables
	envGrpcServer := helper.GetEnvOrDefault("GRPC_REMOTE_SERVER", "localhost:7000")
	envGrpcTls := helper.GetEnvOrDefault("GRPC_TLS", false)
	envGatewayPort := helper.GetEnvOrDefault("GATEWAY_PORT", 8081)

	log.Printf("gRPC Remote Server: %s", envGrpcServer)
	log.Printf("gRPC TLS Enabled: %v", envGrpcTls)
	log.Printf("Gateway Port: %d", envGatewayPort)

	// Setup gRPC Gateway
	serveMux := runtime.NewServeMux()
	gatewayConfig := &gateway.GatewayConfig{
		GrpcRemoteServer: envGrpcServer,
		GrpcTLS:          envGrpcTls,
		ServeMux:         serveMux,
	}

	// Setup graceful shutdown
	shutdown, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Register gRPC services to the gateway
	err = gateway.RegisterHandlerFromEndpoint(shutdown, gatewayConfig)
	if err != nil {
		log.Printf("Failed to register gRPC gateway: %v", err)
		stop()
		return
	}

	r := mux.NewRouter()

	// Register Swagger Handler
	swaggerHandler := handler.NewSwaggerHandler(r)
	swaggerHandler.RegisterRoute()

	// Register GRPC Handler and should be last
	grpcGatewayHandler := handler.NewGrpcGatewayHandler(r, serveMux)
	grpcGatewayHandler.RegisterRoute()

	// Start HTTP server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", envGatewayPort),
		Handler: r,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Starting GRPC-Gateway server on port %s...", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("ListenAndServe error: %v", err)
			stop()
		}
	}()

	<-shutdown.Done()
	log.Println("GRPC-Gateway server is shutting down...")

	// Create a context with timeout for the shutdown process
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("GRPC-Gateway server shutdown failed: %v", err)
		return
	}

	log.Println("GRPC-Gateway server stopped.")
}
