package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func CreateInstance() *mongo.Client {

	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Warning: unable to find .env file.")
	}

	MongoDB := os.Getenv("MONGODB_URI")

	if MongoDB == "" {
		log.Fatal("Unable to load MONGODB_URI environment variable!")
	}

	fmt.Println("MongoDB URI: ", MongoDB) // For debug

	clientOptions := options.Client().ApplyURI(MongoDB)

	client, err := mongo.Connect(clientOptions)

	if err != nil {
		return nil
	}

	return client
}

var Client *mongo.Client = CreateInstance()

func OpenCollection(collectionName string) *mongo.Collection {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: unable to find .env file.")
	}

	databaseName := os.Getenv("DATABASE_NAME")

	fmt.Println("DATABASE_NAME: ", databaseName) // For debug

	collection := Client.Database(databaseName).Collection(collectionName)

	if collection == nil {
		return nil
	}

	return collection
}
