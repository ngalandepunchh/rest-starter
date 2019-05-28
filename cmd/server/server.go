package server

import (
	"log"

	"github.com/hunterpunchh/rest-starter/internal/handlers"
	"github.com/hunterpunchh/rest-starter/internal/router"
	"google.golang.org/grpc"

	statuspb "github.com/hunterpunchh/grpc-starter/pkg/status"
	"go.uber.org/dig"
)

// Run configures, builds, and runs a new server
func Run() {
	c := dig.New()

	_ = c.Provide(func() (statuspb.GrpcStarterClient, error) {
		conn, err := grpc.Dial("grpc-starter:50051", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		// defer conn.Close()
		return statuspb.NewGrpcStarterClient(conn), nil
	})
	_ = c.Provide(handlers.New)
	_ = c.Provide(router.New)

	err := c.Invoke(func(r router.Router) error {
		log.Println("Running REST Server")
		return r.Serve()
	})
	if err != nil {
		panic(err)
	}

	return
}
