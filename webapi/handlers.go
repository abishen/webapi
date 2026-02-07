package webapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "album not found"})
}
func postAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// deleteAlbum deletes an album by ID.
func deleteAlbum(c *gin.Context) {
	id := c.Param("id")
	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "album not found"})
}

func RegisterRoutes(router *gin.Engine) {
	router.GET("/album", getAlbums)
	router.GET("/album/:id", getAlbumByID)
	router.POST("/album", postAlbum)
	router.DELETE("/album/:id", deleteAlbum)
}
