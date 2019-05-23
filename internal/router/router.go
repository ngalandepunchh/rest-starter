package router

import (
	"net/http"

	"github.com/hunterpunchh/starter/internal/handlers"

	"github.com/go-chi/chi"
)

type router struct {
	http.Handler
}

type Router interface {
	Serve() error
}

func New(h handlers.Handler) Router {
	r := chi.NewRouter()

	r.Mount("/", newIndexRouter(h))

	router := &router{
		Handler: r,
	}

	return router
}

func (r *router) Serve() error {
	return http.ListenAndServe(":8080", r)
}
