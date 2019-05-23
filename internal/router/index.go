package router

import (
	"net/http"

	"github.com/hunterpunchh/starter/internal/handlers"

	"github.com/go-chi/chi"
)

func newIndexRouter(h handlers.IndexHandler) http.Handler {
	r := chi.NewRouter()
	r.Get("/status", h.GetStatus)
	return r
}
