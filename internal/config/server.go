package config

import (
	"errors"
	"net"
	"os"
)

const (
	serverHostEnvName = "SERVER_HOST"
	serverPortEnvName = "SERVER_PORT"
)

type serverConfig struct {
	host string
	port string
}

// NewServerConfig returns new server config
func NewServerConfig() (ServerConfig, error) {
	host := os.Getenv(serverHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("server host not found")
	}

	port := os.Getenv(serverPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("server port not found")
	}

	return &serverConfig{
		host: host,
		port: port,
	}, nil
}

// Address returns full address of a server
func (cfg *serverConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
