package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sidmohanty11/gradbook/server/db"
	"github.com/sidmohanty11/gradbook/server/handlers"
	"github.com/sidmohanty11/gradbook/server/routes"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const PORT = ":8000"

func main() {
	// server connection
	app := fiber.New()
	
	app.Use(cors.New(cors.Config{})) //cross-origin-resource-sharing
	app.Use(logger.New())
	
	// db connection
	db, err := db.ConnectSQL()
	
	if err != nil {
		log.Fatalln("Cannot connect to DB!")
	}

	log.Println("Connected to DB at PORT 5432")
	defer db.SQL.Close()
	
	routes.Setup(app)
	handlers.NewRepo(db)
	// server listening port
	app.Listen(PORT)
	
	log.Printf("Listening at PORT%s", PORT)
}
