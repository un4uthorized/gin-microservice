package main 

import (
	"net/http"
	"github.com/gin-gonic/gin"	
)

type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album{
	{ ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}

func CreateAlbums(ctx *gin.Context) {
	var newAlbum album

	if err := ctx.BindJSON(&newAlbum); err != nil {
		return 
	}

	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func DeleteAlbums(ctx *gin.Context) {
	id := ctx.Param("id")

	for i, album := range albums {
		if album.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			break
		}
	}

	ctx.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumByID(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, a := range albums {
		if a.ID == id {
			ctx.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", GetAlbums)
	router.GET("/albums/:id", GetAlbumByID)
	router.POST("/albums", CreateAlbums)
	router.DELETE("/albums/:id", DeleteAlbums)
	router.Run("localhost:8080")
}