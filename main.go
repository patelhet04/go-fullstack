package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        int    `json:"id" bson:"_id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("Hello World")
	// load fibre library
	app := fiber.New()
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading env file")
	}

	MONGO_URI := os.Getenv("MONGO_URI")
	PORT := os.Getenv("PORT")
	clientOptions := options.Client().ApplyURI(MONGO_URI)

	// Initializes the mongodb connection
	// context.Background(): This provides a default context. In Go, contexts are used to manage
	// deadlines, cancellation signals, and other request-scoped values across API boundaries.
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Ping method is used to ensure that MongoDB connection is reachable and the connection is alive
	// This step is crucial to confirm that the connection to the server has been successfully
	// established and the server is responsive
	err = client.Ping(context.Background(), nil)
	// If the ping fails, it logs the error and exits the program.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	app.Listen(":" + PORT)

}
