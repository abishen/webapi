package webapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getAlbums retrieves all albums
func getAlbums(c *gin.Context) {
	albumsMutex.RLock()
	defer albumsMutex.RUnlock()

	// Convert map to slice for JSON response
	result := make([]album, 0, len(albums))
	for _, a := range albums {
		result = append(result, a)
	}

	c.JSON(http.StatusOK, result)
}

// getAlbumByID retrieves a specific album by ID
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	albumsMutex.RLock()
	a, found := albums[id]
	albumsMutex.RUnlock()

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
		return
	}

	c.JSON(http.StatusOK, a)
}

// postAlbum creates a new album
func postAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate required fields
	if newAlbum.ID == "" || newAlbum.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id and title are required"})
		return
	}

	albumsMutex.Lock()
	defer albumsMutex.Unlock()

	// Check if album already exists
	if _, exists := albums[newAlbum.ID]; exists {
		c.JSON(http.StatusConflict, gin.H{"error": "album with this id already exists"})
		return
	}

	albums[newAlbum.ID] = newAlbum
	c.JSON(http.StatusCreated, newAlbum)
}

// deleteAlbum deletes an album by ID
func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	albumsMutex.Lock()
	defer albumsMutex.Unlock()

	if _, found := albums[id]; !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
		return
	}

	delete(albums, id)
	c.JSON(http.StatusOK, gin.H{"message": "album deleted successfully"})
}

// RegisterRoutes sets up all API routes
func RegisterRoutes(router *gin.Engine) {
	router.GET("/album", getAlbums)
	router.GET("/album/:id", getAlbumByID)
	router.POST("/album", postAlbum)
	router.DELETE("/album/:id", deleteAlbum)
}
