package main

import (
	// Internal
	"github.com/hunterpunchh/rest-starter/internal/config"
	"github.com/hunterpunchh/rest-starter/internal/handlers"
	"github.com/hunterpunchh/rest-starter/internal/logger"
	"github.com/hunterpunchh/rest-starter/internal/router"

	// External
	statuspb "github.com/hunterpunchh/grpc-starter/pkg/status"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Run configures, builds, and runs a new server
func main() {
	c := dig.New()

	_ = c.Provide(config.New)

	_ = c.Provide(logger.New)

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

	err := c.Invoke(func(cfg *config.Config, r router.Router, l *zap.SugaredLogger) error {
		defer l.Sync()
		return r.Serve(cfg.Port)
	})
	if err != nil {
		panic(err)
	}

	return
}
