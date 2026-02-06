package webapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func RegisterRoutes(router *gin.Engine) {
	router.GET("/album", getAlbums)
}
