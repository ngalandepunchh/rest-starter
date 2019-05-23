package handlers

import (
	"net/http"

	"github.com/hunterpunchh/starter/internal/responses"
)

type indexHandler struct {
}

type IndexHandler interface {
	GetStatus(w http.ResponseWriter, r *http.Request)
}

func NewIndexHandler() IndexHandler {
	return &indexHandler{}
}

func (h *indexHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	resp := responses.Status{Status: responses.StatusOK}
	_ = WriteJSON(w, resp, http.StatusOK)
	return
}
