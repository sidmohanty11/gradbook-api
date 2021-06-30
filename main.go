package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sidmohanty11/gradbook/server/db"
	"github.com/sidmohanty11/gradbook/server/routes"
)

func main() {
	// server connection
	app := fiber.New()
	routes.Setup(app)

	app.Use(cors.New(cors.Config{})) //cross-origin-resource-sharing

	// db connection
	conn := db.Connect()
	defer conn.Close(context.Background())

	// server listening port
	app.Listen(":8000")
}
