package handlers

import (
	"net/http"
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
	resp := struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	}

	_ = WriteJSON(w, resp, http.StatusOK)

	return
}
