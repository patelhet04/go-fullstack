package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	// Here instead of int we have to use primitive.ObjectID, as mongoDB as its own _id which is of type ObejctId
	// Furthermore we are using omitempty, as the default value it takes is 0, which will create a _id = ObjectId('000000000000')
	// So we will just skip the false or default value which is 0
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("Hello World")
	// load fibre library
	app := fiber.New()

	// cors
	app.Use(cors.New(cors.Config{AllowOrigins: "http://localhost:5173", AllowHeaders: "Origin, Content-Type, Accept"}))
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

	// defer postpones the excecution of a function on its R.H.S
	// until the execution of surounding function is completed
	defer client.Disconnect(context.Background())

	fmt.Println("Connected to MongoDB")

	collection = client.Database("Go-WebApp").Collection("todos")

	// API endpoints
	app.Get("/api/todos", getTodos)
	app.Post("/api/todos", createTodo)
	app.Patch("/api/todos/:id", updateTodos)
	app.Delete("/api/todos/:id", deleteTodos)
	log.Fatal(app.Listen("0.0.0.0:" + PORT))
}

func getTodos(c *fiber.Ctx) error {
	var todos []Todo
	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	// Iterating through the documents
	for cursor.Next(context.Background()) {
		var todo Todo
		// Decode function is used to convert a BSON representation of a MongoDB document into a Go struct.
		// Decode takes in the address of struc as parameter, that's why we have used '&' operator
		if err := cursor.Decode(&todo); err != nil {
			return err
		}
		todos = append(todos, todo)
	}
	return c.JSON(todos)
}

func createTodo(c *fiber.Ctx) error {
	// This is going to be a pointer
	// Default values {id:0, completed:false, body:""}
	todo := new(Todo)

	// Body parser binds the json request into Todo struct
	if err := c.BodyParser(todo); err != nil {
		return err
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"msg": "Request body cannot be empty"})
	}

	insertResult, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}

	todo.ID = insertResult.InsertedID.(primitive.ObjectID)
	return c.Status(200).JSON(todo)
}

func updateTodos(c *fiber.Ctx) error {
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.Status(404).JSON(fiber.Map{"msg": "ID cannot be found"})
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"completed": true}}
	_, err = collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{"msg": "success"})
}

func deleteTodos(c *fiber.Ctx) error {
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.Status(404).JSON(fiber.Map{"msg": "ID cannot be found"})
	}

	filter := bson.M{"_id": objectId}
	_, err = collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"msg": "sucess"})
}
