package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/nathalia-nobrega/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
)

func main() {

	router := gin.Default()

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.String(200, "LOve will Be REvealed!")
	})

	router.GET("/movies", controllers.GetMovies())

	router.GET("/unique-movie/:imdb_id", controllers.GetUniqueMovie())

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}

}
