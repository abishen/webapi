package main

import (
	webapi "github.com/abishen/webapi/webapi"

	"github.com/gin-gonic/gin"
)

func main() {
	config := webapi.LoadConfig()

	router := gin.Default()
	webapi.RegisterRoutes(router)
	router.Run(config.ServerAddr)

}
