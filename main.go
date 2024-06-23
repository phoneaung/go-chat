package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/phoneaung/go-chat/handlers"
)

func main() {
	// Create views engine
	viewsEngine := html.New("./views", ".html")

	// Start new fiber instance
	app := fiber.New(fiber.Config{
		Views: viewsEngine,
	})

	// Create new App Handler
	appHandler := handlers.NewAppHandler()

	// Add app handler route
	app.Get("/", appHandler.HandleGetIndex)

	// Static route and directory
	app.Static("/static/", "./static")

	// Create a "ping" handler to test the server
	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome to fiber!")
	})

	// Start the http server
	app.Listen(":3000")
	fmt.Println("Hello worlds!")
}
