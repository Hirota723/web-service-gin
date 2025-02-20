package controller

import (
	"example/web-service-gin/model"
	"example/web-service-gin/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func GetAlbumsController(c *gin.Context) {
	albums := repository.GetAlbums()
	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByIDController(c *gin.Context) {
	id := c.Param("id")
	album, found := repository.GetAlbumByID(id)

	if !found {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}


// postAlbums adds an album from JSON received in the request body.
func PostAlbumsController(c *gin.Context) {
	var newAlbum model.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	repository.AddAlbum(newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}