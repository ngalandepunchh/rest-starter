package router

import (
	"net/http"

	"github.com/hunterpunchh/rest-starter/internal/handlers"

	"github.com/go-chi/chi"
)

func newIndexRouter(h handlers.IndexHandler) http.Handler {
	r := chi.NewRouter()
	r.Get("/status", h.GetStatus)
	r.Get("/status", h.GetHealthz)
	return r
}
