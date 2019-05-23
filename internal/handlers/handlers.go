package handlers

import (
	"encoding/json"
	"net/http"
)

type handler struct {
	IndexHandler
}

// Handler defines an interface for interacting with all handlers
type Handler interface {
	IndexHandler
}

// New creates and returns a new Handler
func New() Handler {
	return &handler{
		IndexHandler: NewIndexHandler(),
	}
}

// WriteJSON sets the status code and writes data to the given ResponseWriter
func WriteJSON(w http.ResponseWriter, data interface{}, code int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(data)
}
