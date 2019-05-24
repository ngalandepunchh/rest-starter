package handlers

import (
	"net/http"

	"github.com/hunterpunchh/rest-starter/internal/responses"
)

type indexHandler struct {
}

// IndexHandler defines an interface for interacting with index routes
type IndexHandler interface {
	GetStatus(w http.ResponseWriter, r *http.Request)
	GetHealthz(w http.ResponseWriter, r *http.Request)
}

// newIndexHandler creates and returns a new IndexHandler
func newIndexHandler() IndexHandler {
	return &indexHandler{}
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
	resp := responses.Status{Status: responses.StatusOK}
	_ = WriteJSON(w, resp, http.StatusOK)
	return
}
