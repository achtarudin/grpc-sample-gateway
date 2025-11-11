package handler

import (
	"bytes"
	handlerPort "grpc-sample-gateway/internal/port/handler-port"
	"net/http"
	"time"

	asset "github.com/achtarudin/grpc-sample/openapiv2"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type swaggerHandler struct {
	mux *mux.Router
}

func NewSwaggerHandler(r *mux.Router) handlerPort.HandlerPort {
	return &swaggerHandler{
		mux: r,
	}
}

func (s *swaggerHandler) RegisterRoute() {

	// Serve swagger.json
	s.mux.HandleFunc("/swagger.json", s.serveSwaggerFile).Methods(http.MethodGet)

	// Serve Swagger UI
	s.mux.PathPrefix("/doc/").Handler(s.serveSwaggerUi())
}

func (s *swaggerHandler) serveSwaggerFile(w http.ResponseWriter, r *http.Request) {
	fileBytes, err := asset.SwaggerFS.ReadFile("openapiv2.swagger.json")
	if err != nil {
		http.Error(w, "swagger file not embedded", http.StatusInternalServerError)
		return
	}
	http.ServeContent(w, r, "swagger.json", time.Now(), bytes.NewReader(fileBytes))
}

func (s *swaggerHandler) serveSwaggerUi() http.HandlerFunc {
	return httpSwagger.Handler(
		httpSwagger.PersistAuthorization(true),
		httpSwagger.UIConfig(map[string]string{
			"defaultModelsExpandDepth": "-1",
		}),
		httpSwagger.URL("/swagger.json"),
	)
}
