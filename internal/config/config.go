package config

import (
	"log"
	"os"
)

type Config struct {
	Port               string
	GRPCStarterAddress string
}

func New() *Config {
	log.Println(os.Environ())
	port, _ := os.LookupEnv("PORT")
	grpcStarterAddress, _ := os.LookupEnv("GRPC_STARTER_ADDRESS")

	cfg := &Config{
		Port:               port,
		GRPCStarterAddress: grpcStarterAddress,
	}

	log.Println(cfg)

	return cfg
}
