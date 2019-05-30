package server

import (
	// Internal
	"github.com/hunterpunchh/rest-starter/internal/config"
	"github.com/hunterpunchh/rest-starter/internal/handlers"
	"github.com/hunterpunchh/rest-starter/internal/router"

	// External
	statuspb "github.com/hunterpunchh/grpc-starter/pkg/status"
	"go.uber.org/dig"
	"google.golang.org/grpc"
)

// Run configures, builds, and runs a new server
func Run() {
	c := dig.New()

	_ = c.Provide(config.New)

	_ = c.Provide(func(cfg *config.Config) (statuspb.GrpcStarterClient, error) {
		conn, err := grpc.Dial(cfg.GRPCStarterAddress, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		// defer conn.Close()
		return statuspb.NewGrpcStarterClient(conn), nil
	})

	_ = c.Provide(handlers.New)
	_ = c.Provide(router.New)

	err := c.Invoke(func(cfg *config.Config, r router.Router) error {
		return r.Serve(cfg.Port)
	})
	if err != nil {
		panic(err)
	}

	return
}
