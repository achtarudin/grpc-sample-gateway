package handler

import (
	handlerPort "grpc-sample-gateway/internal/port/handler-port"
	"net/http"

	"github.com/gorilla/mux"
)

type grpcGatewayHandler struct {
	mux       *mux.Router
	serverMux http.Handler
}

func NewGrpcGatewayHandler(r *mux.Router, serverMux http.Handler) handlerPort.HandlerPort {
	return &grpcGatewayHandler{
		mux:       r,
		serverMux: serverMux,
	}
}

func (s *grpcGatewayHandler) RegisterRoute() {
	// Serve gRPC-Gateway
	s.mux.PathPrefix("/").Handler(s.serverMux)
}
