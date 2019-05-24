package server

import (
	"github.com/hunterpunchh/rest-starter/internal/handlers"
	"github.com/hunterpunchh/rest-starter/internal/router"

	"go.uber.org/dig"
)

// Run configures, builds, and runs a new server
func Run() {
	c := dig.New()

	_ = c.Provide(handlers.New)
	_ = c.Provide(router.New)

	err := c.Invoke(func(r router.Router) error {
		return r.Serve()
	})
	if err != nil {
		panic(err)
	}

	return
}
