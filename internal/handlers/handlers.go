package handlers

import (
	"encoding/json"
	"net/http"
)

type handler struct {
	IndexHandler
}

type Handler interface {
	IndexHandler
}

func New() Handler {
	return &handler{
		IndexHandler: NewIndexHandler(),
	}
}

func WriteJSON(w http.ResponseWriter, data interface{}, code int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(data)
}
