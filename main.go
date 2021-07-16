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
		AllowHeaders:     "Origin, Content-Type, Accept, Accept-Language, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
	}))

	app.Use(logger.New())

	// csrf protection...
	// app.Use(csrf.New(csrf.Config{
	// 	KeyLookup:      "header:X-CSRF-Token",
	// 	CookieName:     "csrf_",
	// 	CookieSameSite: "Strict",
	// 	Expiration:     1 * time.Hour,
	// 	KeyGenerator:   utils.UUID,
	// }))

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
	if os.Getenv("APP_ENV") != "production" {
		_ = godotenv.Load()
	}

	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}
