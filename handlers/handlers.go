package handlers

import (
	"database/sql"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sidmohanty11/gradbook/server/db"
)

// creates a global postgresdb instance.
var Psql *sql.DB

//creates a new repository, gets the db conn from the main.go file
func NewRepo(db *db.DB) {
	Psql = db.SQL
}

func ValidToken(c *fiber.Ctx) (*jwt.StandardClaims, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	cookie := c.Cookies("sid")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	// user is not logged in.
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	return claims, nil
}
