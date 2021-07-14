package handlers

import (
	"database/sql"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/sidmohanty11/gradbook/server/db"
)

// creates a global postgresdb instance.
var Psql *sql.DB

//creates a new repository, gets the db conn from the main.go file
func NewRepo(db *db.DB) {
	Psql = db.SQL
}

func ValidToken(c *fiber.Ctx) (*jwt.StandardClaims, error) {
	cookie := c.Cookies("sid")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	// user is not logged in.
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	return claims, nil
}
