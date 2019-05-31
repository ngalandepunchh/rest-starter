package config

import (
	"os"
)

// Config is a struct containing necessary configuration values
type Config struct {
	Port               string
	GRPCStarterAddress string
}

// New creates a config from the defined environment variables
func New() *Config {
	port, _ := os.LookupEnv("PORT")
	grpcStarterAddress, _ := os.LookupEnv("GRPC_STARTER_ADDRESS")

	cfg := &Config{
		Port:               port,
		GRPCStarterAddress: grpcStarterAddress,
	}

	return cfg
}
