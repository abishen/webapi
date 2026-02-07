package webapi

import (
	"os"
	"strings"
)

// Config holds application configuration
type Config struct {
	ServerAddr string
	GinMode    string
	Port       string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	addr := os.Getenv("SERVER_ADDR")
	if addr == "" {
		addr = "localhost:8080"
	}

	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = "debug"
	}
	// Validate gin mode
	if !isValidGinMode(mode) {
		mode = "debug"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		ServerAddr: addr,
		GinMode:    mode,
		Port:       port,
	}
}

// isValidGinMode validates gin mode value
func isValidGinMode(mode string) bool {
	validModes := []string{"debug", "release", "test"}
	mode = strings.ToLower(mode)
	for _, valid := range validModes {
		if mode == valid {
			return true
		}
	}
	return false
}
