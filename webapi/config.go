package webapi

import (
	"os"
)

type Config struct {
	ServerAddr string
}

func LoadConfig() *Config {
	addr := os.Getenv("SERVER_ADDR")
	if addr == "" {
		addr = "localhost:8080"
	}

	return &Config{ServerAddr: addr}
}
