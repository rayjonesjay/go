package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

/*
getAlbums takes a gin context as a parameter
gin.Context is the most important part of gin
it carries request details
validates and serializes json ...etc

// IndentedJSON serializes the album slice and adds it to the response
*/
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getUserById(c *gin.Context) {
	r := c.Request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(body)
	i, _ := strconv.Atoi(string(body))

	fmt.Println(i)

	for _, cc := range albums {
		if cc.ID == string(body) {
			c.IndentedJSON(http.StatusOK, c)
			return
		}
	}

}

func postAlbum(c *gin.Context) {
	var newAlbum album

	err := c.BindJSON(&newAlbum)
	if err != nil {
		return
	}
	albums = append(albums, newAlbum)
	// you can use Context.IndentedJSON in development but not preferred during production
	// since it consumes alot of CPU when prettifying or formating the json object
	// use Context.JSON
	c.JSON(http.StatusCreated, newAlbum)
}

func GetUserById(c *gin.Context) {
	// use Context.Param to retrieve the id path parameter from the URL
	id := c.Param("id")
	for _, cc := range albums {
		if cc.ID == id {
			c.IndentedJSON(http.StatusOK, cc)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.GET("/albums/:id", GetUserById)
	router.Run("localhost:8080")
	f()
}

func f(a ...int) {
	for _, c := range a {
		fmt.Println(c)
	}
}
