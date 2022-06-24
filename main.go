package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create new Fiber instance
	app := fiber.New()

	// Create new GET route on path "/hello"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Start server on http://localhost:3000
	log.Fatal(app.Listen(":3000"))
}
