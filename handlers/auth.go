package handlers

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sidmohanty11/gradbook/server/helpers"
	"golang.org/x/crypto/bcrypt"
)

// the register handler which takes in the inputs and checks whether everything is valid,
// hashes the password thats given through bcrypt
func Register(c *fiber.Ctx) error {
	if os.Getenv("APP_ENV") != "production" {
		_ = godotenv.Load()
	}

	type RegisterInput struct {
		Username          string `json:"username"`
		Password          string `json:"password"`
		Email             string `json:"email"`
		ImageURL          string `json:"image_url"`
		TheSecretPasscode string `json:"passcode"`
	}
	var input RegisterInput

	err := c.BodyParser(&input)

	if err != nil || input.TheSecretPasscode != os.Getenv("SECRET_PASSCODE") {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	identity := input.Username
	pass := input.Password
	imgurl := input.ImageURL
	email := input.Email

	hashedpass, err := helpers.HashPassword(pass)

	if err != nil {
		return err
	}

	row := Psql.QueryRow("insert into users (username, image_url, email) values ($1,$2,$3) returning id;", identity, imgurl, email)

	var id int

	row.Scan(&id)

	_, err = Psql.Exec("insert into login (user_id, username, hash) values ($1,$2,$3);", id, identity, hashedpass)

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success"})
}

// checks if the user is present in the db,
// returns a signed jwt token
func Login(c *fiber.Ctx) error {
	if os.Getenv("APP_ENV") != "production" {
		_ = godotenv.Load()
	}

	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	identity := input.Username
	pass := input.Password

	row := Psql.QueryRow("SELECT username, hash FROM login WHERE username = $1;", identity)

	var dbUsername string
	var dbPassword string

	if err := row.Scan(&dbUsername, &dbPassword); err != nil {
		return err
	}

	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(pass))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return c.SendStatus(fiber.StatusUnauthorized)
	} else if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if identity != dbUsername {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = identity
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "sid",
		Value:    t,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{"status": "success", "username": dbUsername})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "sid",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "logout successful",
	})
}
