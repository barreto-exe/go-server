package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "2", Title: "Blurryface", Artist: "Twenty One Pilots", Price: 20.99},
	{ID: "3", Title: "Trench", Artist: "Twenty One Pilots", Price: 20.99},
	{ID: "4", Title: "Vessel", Artist: "Twenty One Pilots", Price: 20.99},
	{ID: "5", Title: "Regional at Best", Artist: "Twenty One Pilots", Price: 20.99},
	{ID: "6", Title: "Twenty One Pilots", Artist: "Twenty One Pilots", Price: 20.99},
}

func getAlbumns(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbumns)
	router.POST("/albums", postAlbums)

	router.Run("localhost:3000")
}
