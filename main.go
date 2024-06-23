package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/phoneaung/go-chat/handlers"
)

func main() {

	// Start new fiber instance
	app := fiber.New()

	// Create a "ping" handler to test the server
	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome to fiber")
	})

	// create new App Handler
	appHandler := handlers.NewAppHandler()

	// Add appHandler routes
	app.Get("/", appHandler.HandleGetIndex)

	// create new webscoket
	server := NewWebSocket()
	app.Get("/ws", websocket.New(func(ctx *websocket.Conn) {
		server.HandleWebSocket(ctx)
	}))

	go server.HandleMessages()

	// Start the http server
	app.Listen(":3000")
}
