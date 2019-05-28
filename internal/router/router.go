package router

import (
	"net/http"

	"github.com/hunterpunchh/rest-starter/internal/handlers"

	"github.com/go-chi/chi"
)

type router struct {
	http.Handler
}

// Router defines an interface for interacting with the router
// Router implements the http.Handler interface
type Router interface {
	http.Handler
	Serve() error
}

// New creates a Router from a given handler.Handler
func New(h handlers.Handler) Router {
	r := chi.NewRouter()

	r.Mount("/", newIndexRouter(h))

	r.NotFound(h.NotFound)
	r.MethodNotAllowed(h.MethodNotAllowed)

	router := &router{
		Handler: r,
	}

	return router
}

// Serve listens on port :8080 and serves the given requests
func (r *router) Serve() error {
	return http.ListenAndServe(":8080", r)
}
