package handlers

import (
	"encoding/json"
	"net/http"

	statuspb "github.com/hunterpunchh/grpc-starter/pkg/status"
	"go.uber.org/zap"
)

type handler struct {
	logger *zap.SugaredLogger
	IndexHandler
}

// Handler defines an interface for interacting with all handlers
type Handler interface {
	IndexHandler
	NotFound(w http.ResponseWriter, r *http.Request)
	MethodNotAllowed(w http.ResponseWriter, r *http.Request)
}

// New creates and returns a new Handler
func New(grpcStarterClient statuspb.GrpcStarterClient, l *zap.SugaredLogger) Handler {
	handler := &handler{
		logger:       l,
		IndexHandler: newIndexHandler(grpcStarterClient, l),
	}

	return handler
}

func (h *handler) NotFound(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, nil, http.StatusNotFound)
}

func (h *handler) MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, nil, http.StatusMethodNotAllowed)
}

// WriteJSON sets the status code and writes data to the given ResponseWriter
func WriteJSON(w http.ResponseWriter, data interface{}, code int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(data)
}
