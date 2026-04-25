package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album struct declaration
// store album data in memory
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// slice of albums to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// sets up association where getAlbums handles requests to the /albums endpoint path
func main() {
	router := gin.Default()          // initialize Gin using Default
	router.GET("/albums", getAlbums) // use GET function to associate the GET HTTP method and /albums path with handler function

	router.Run("localhost:8080") // use the Run function to attach the route to an http.Server and start the server
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) { // gin.Context very important part of Gin it carries requestt details, validates and serializes JSON, n more
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// call BindJSON to bind the recieved JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
