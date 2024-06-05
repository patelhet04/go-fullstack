package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Hello world")

	todos := []Todo{}

	app := fiber.New()
	// To fetch todo items
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"name": "John"})
	})

	// To create a Todo
	app.Post("/api/todo", func(c *fiber.Ctx) error {
		todo := &Todo{}
		// In Go, whenver using a pointer always check for errors
		// BodyParser binds the request body to the struct
		if err := c.BodyParser(todo); err != nil {
			return err
		}
		if todo.Body == "" {
			c.Status(400).JSON(fiber.Map{"msg": "Body cannot be empty"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)
		return c.Status(201).JSON(todo)
	})

	// To update a Todo
	app.Patch("/api/todo/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = !todos[i].Completed
				return c.Status(200).JSON(todos[i])
			}
		}
		return c.Status(400).JSON(fiber.Map{"msg": "No todo found!"})

	})

	log.Fatal(app.Listen(":4000"))
}
