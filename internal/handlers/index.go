package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/hunterpunchh/rest-starter/internal/responses"

	statuspb "github.com/hunterpunchh/grpc-starter/pkg/status"
)

type indexHandler struct {
	statuspb.GrpcStarterClient
}

// IndexHandler defines an interface for interacting with index routes
type IndexHandler interface {
	GetStatus(w http.ResponseWriter, r *http.Request)
	GetHealthz(w http.ResponseWriter, r *http.Request)
}

// newIndexHandler creates and returns a new IndexHandler
func newIndexHandler(client statuspb.GrpcStarterClient) IndexHandler {
	return &indexHandler{
		GrpcStarterClient: client,
	}
}

// GetStatus returns a Status response with the current status of the service
// If the service is running "ok" is returned with status 200
func (h *indexHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	resp := responses.Status{Status: responses.StatusOK}
	_ = WriteJSON(w, resp, http.StatusOK)
	return
}

// GetHealthz returns a Status response with the current health of the service
func (h *indexHandler) GetHealthz(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := h.GrpcStarterClient.Status(ctx, &statuspb.StatusRequest{})
	if err != nil {
		log.Println(err)
		_ = WriteJSON(w, nil, http.StatusInternalServerError)
	}
	_ = WriteJSON(w, resp, http.StatusOK)
	return
}
