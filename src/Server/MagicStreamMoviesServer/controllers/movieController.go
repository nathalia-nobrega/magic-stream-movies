package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nathalia-nobrega/MagicStreamMovies/Server/MagicStreamMoviesServer/database"
	"github.com/nathalia-nobrega/MagicStreamMovies/Server/MagicStreamMoviesServer/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var movieCollection *mongo.Collection = database.OpenCollection("movies")

func GetMovies() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 100*time.Second) // memory management
		defer cancel()

		var movies []models.Movie
		cursor, err := movieCollection.Find(context, bson.M{})

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies."})
		}
		defer cursor.Close(context) // memory management

		if err = cursor.All(context, &movies); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies."})
		}

		ctx.JSON(http.StatusOK, movies)
	}

}

func GetUniqueMovie() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context, cancel := context.WithTimeout(context.Background(), 100*time.Second) // memory management i assume
		defer cancel()

		movieID := ctx.Param("imdb_id")

		// exception handling
		if movieID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Movie ID is required!"})
		}

		var uniqueMovie models.Movie

		err := movieCollection.FindOne(context, bson.M{"imdb_id": movieID}).Decode(&uniqueMovie)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Movie not found."})
			return
		}

		ctx.JSON(http.StatusOK, uniqueMovie)
	}
}
