package starter

import (
	"github.com/hunterpunchh/starter/internal/handlers"
	"github.com/hunterpunchh/starter/internal/router"

	"go.uber.org/dig"
)

func main() {
	c := dig.New()

	_ = c.Provide(handlers.New)
	_ = c.Provide(router.New)

	_ = c.Invoke(func(r router.Router) {
		r.Serve()
	})

	return
}
