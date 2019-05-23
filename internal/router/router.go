package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/hunterpunchh/starter/internal/handlers"
)

type router struct {
	http.Handler
}

type Router interface {
}

func New(h handlers.Handler) Router {
	r := chi.NewRouter()

	r.Mount("/", newIndexRouter(h))

	router := &router{
		Handler: r,
	}

	return router
}
