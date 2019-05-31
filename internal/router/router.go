package router

import (
	"net/http"

	// Internal
	"github.com/hunterpunchh/rest-starter/internal/handlers"
	"github.com/hunterpunchh/rest-starter/internal/middleware"

	// External
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type router struct {
	http.Handler
	logger *zap.SugaredLogger
}

// Router defines an interface for interacting with the router
// Router implements the http.Handler interface
type Router interface {
	http.Handler
	Serve(string) error
}

// New creates a Router from a given handler.Handler
func New(h handlers.Handler, l *zap.SugaredLogger) Router {
	r := chi.NewRouter()

	r.Mount("/", newIndexRouter(h))

	r.NotFound(h.NotFound)
	r.MethodNotAllowed(h.MethodNotAllowed)

	router := &router{
		Handler: r,
		logger:  l,
	}

	router.Handler = middleware.Log(router.Handler, router.logger)

	return router
}

// Serve listens and serves requests on the provided port
func (r *router) Serve(port string) error {
	return http.ListenAndServe(":"+port, r)
}
