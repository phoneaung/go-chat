package main

import "github.com/gofiber/fiber/v2"

func main() {

	// Start new fiber instance
	app := fiber.New()

	// Create a "ping" handler to test the server
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to fiber!")
	})

	// Start the http server
	app.Listen(":3000")
}
