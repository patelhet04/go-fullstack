package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type TodoInMemory struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func mainInMemory() {
	fmt.Println("Hello world")

	todoInMemorys := []TodoInMemory{}

	app := fiber.New()

	// Load godotenv
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading env file")
	}

	PORT := os.Getenv("PORT")
	// To fetch TodoInMemory items
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"name": "John"})
	})

	// To create a TodoInMemory
	app.Post("/api/TodoInMemory", func(c *fiber.Ctx) error {
		TodoInMemory := &TodoInMemory{}
		// In Go, whenver using a pointer always check for errors
		// BodyParser binds the request body to the struct
		if err := c.BodyParser(TodoInMemory); err != nil {
			return err
		}
		if TodoInMemory.Body == "" {
			return c.Status(400).JSON(fiber.Map{"msg": "Body cannot be empty"})
		}

		TodoInMemory.ID = len(todoInMemorys) + 1
		todoInMemorys = append(todoInMemorys, *TodoInMemory)
		return c.Status(201).JSON(TodoInMemory)
	})

	// To update a TodoInMemory
	app.Patch("/api/TodoInMemory/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, TodoInMemory := range todoInMemorys {
			// Formats the value and returns a string
			if fmt.Sprint(TodoInMemory.ID) == id {
				todoInMemorys[i].Completed = !todoInMemorys[i].Completed
				return c.Status(200).JSON(todoInMemorys[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{"msg": "No TodoInMemory found!"})
	})

	//  Delete a TodoInMemory
	app.Delete("/api/TodoInMemory/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, TodoInMemory := range todoInMemorys {
			if fmt.Sprint(TodoInMemory.ID) == id {
				todoInMemorys = append(todoInMemorys[:i], todoInMemorys[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"msg": "Success"})
			}
		}
		return c.Status(404).JSON(fiber.Map{"msg": "No TodoInMemory found!"})
	})
	log.Fatal(app.Listen(":" + PORT))
}
