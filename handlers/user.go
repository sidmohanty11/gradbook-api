package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sidmohanty11/gradbook/server/models"
)

// finds an user by id, returns data to the frontend as json
func UserById(c *fiber.Ctx) error {
	paramId := c.Params("id")
	row := Psql.QueryRow("SELECT id, username, email, image_url, created_on, last_login FROM users WHERE id = $1;", paramId)

	user := models.User{}
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.ImageURL, &user.CreatedOn, &user.LastLogin); err != nil {
		return err
	}

	claims, err := ValidToken(c)

	if err != nil {
		return c.JSON(fiber.Map{"status": "unauthorized"})
	}

	return c.JSON(fiber.Map{"user": user, "claims": claims})
}

// finds an user by username, returns data to the frontend as json
func UserByUsername(c *fiber.Ctx) error {
	username := c.Params("username")

	row := Psql.QueryRow("SELECT id, username, email, image_url, created_on, last_login FROM users WHERE username = $1;", username)

	user := models.User{}
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.ImageURL, &user.CreatedOn, &user.LastLogin); err != nil {
		return err
	}

	claims, err := ValidToken(c)

	if err != nil {
		return c.JSON(fiber.Map{"status": "unauthorized"})
	}

	return c.JSON(fiber.Map{"user": user, "claims": claims})
}

// deletes a particular user from the db ...
func DeleteUser(c *fiber.Ctx) error {
	paramId := c.Params("id")

	_, err := Psql.Exec("DELETE FROM users WHERE id = $1;", paramId)
	if err != nil {
		return err
	}

	return c.JSON("Deleted")
}
