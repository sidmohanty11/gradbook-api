package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/sidmohanty11/gradbook/server/db"
	"github.com/sidmohanty11/gradbook/server/handlers"
	"github.com/sidmohanty11/gradbook/server/routes"
)

func main() {
	// server connection
	app := fiber.New()

	PORT, err := getPort()

	if err != nil {
		log.Fatal(err)
	}

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	})) //cross-origin-resource-sharing
	app.Use(logger.New())

	// db connection
	db, err := db.ConnectSQL()

	if err != nil {
		log.Fatalln("Cannot connect to DB!")
	}

	fmt.Println("Connected to DB at PORT 5432")
	defer db.SQL.Close()

	routes.Setup(app)
	handlers.NewRepo(db)
	// server listening port
	app.Listen(PORT)

	fmt.Printf("Listening at PORT%s", PORT)
}

func getPort() (string, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}
