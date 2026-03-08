package config

import (
	"log"
	"net"

	"github.com/caarlos0/env/v11"
)

// Config is a struct that contains app configuration
type Config struct {
	LoggerLevel string `env:"LOGGER_LEVEL,required"`
	PostgresDSN string `env:"PG_DSN,required"`
	ServerHost  string `env:"SERVER_HOST,required"`
	ServerPort  string `env:"SERVER_PORT,required"`
}

// MustLoad parses the environment variables and returns the config
func MustLoad() *Config {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	return cfg
}

// ServerAddr returns the server address by combining the server host and port.
func (cfg *Config) ServerAddr() string {
	return net.JoinHostPort(cfg.ServerHost, cfg.ServerPort)
}
