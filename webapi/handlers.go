package webapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var albumStore AlbumStore

// SetAlbumStore sets the album store to use for database operations
func SetAlbumStore(store AlbumStore) {
	albumStore = store
}

// getAlbums retrieves all albums from the database
func getAlbums(c *gin.Context) {
	albums, err := albumStore.GetAllAlbums()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve albums"})
		return
	}

	c.JSON(http.StatusOK, albums)
}

// getAlbumByID retrieves a specific album by ID from the database
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	album, err := albumStore.GetAlbumByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve album"})
		return
	}

	if album == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
		return
	}

	c.JSON(http.StatusOK, album)
}

// postAlbum creates a new album in the database
func postAlbum(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate required fields
	if newAlbum.ID == "" || newAlbum.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id and title are required"})
		return
	}

	// Check if album already exists
	existing, err := albumStore.GetAlbumByID(newAlbum.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check album"})
		return
	}

	if existing != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "album with this id already exists"})
		return
	}

	// Create the album in the database
	if err := albumStore.CreateAlbum(&newAlbum); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create album"})
		return
	}

	c.JSON(http.StatusCreated, newAlbum)
}

// deleteAlbum deletes an album by ID from the database
func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	// Check if album exists
	album, err := albumStore.GetAlbumByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve album"})
		return
	}

	if album == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
		return
	}

	// Delete the album from the database
	if err := albumStore.DeleteAlbum(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete album"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "album deleted successfully"})
}

// RegisterRoutes sets up all API routes
func RegisterRoutes(router *gin.Engine) {
	router.GET("/album", getAlbums)
	router.GET("/album/:id", getAlbumByID)
	router.POST("/album", postAlbum)
	router.DELETE("/album/:id", deleteAlbum)
}
