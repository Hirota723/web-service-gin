package repository

import (
	"example/web-service-gin/model"
)

// albums slice to seed record album data.
var albums = []model.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func GetAlbums() []model.Album {
	return albums
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(id string) (model.Album, bool) {
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			return a, true
		}
	}
	return model.Album{}, false
}

// postAlbums adds an album from JSON received in the request body.
func AddAlbum(newAlbum model.Album) {
	albums = append(albums, newAlbum)
}
