package main

import (
	"log"

	webapi "github.com/abishen/webapi/webapi"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config := webapi.LoadConfig()

	// Set gin mode from config
	gin.SetMode(config.GinMode)

	// Create router
	router := gin.Default()

	// Register all routes
	webapi.RegisterRoutes(router)

	// Start server with error handling
	log.Printf("Starting server on %s\n", config.ServerAddr)
	if err := router.Run(config.ServerAddr); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
