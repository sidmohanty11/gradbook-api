package handlers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/sidmohanty11/gradbook/server/db"
	"github.com/sidmohanty11/gradbook/server/models"
)

var psql *sql.DB

//creates a new repository
func NewRepo(db *db.DB) {
	psql = db.SQL
} 

func Register(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func Login(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func User(c *fiber.Ctx) error {
	rows, err := psql.Query("SELECT id, username, email, hash, image_url, created_on FROM users;")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		defer rows.Close()
		result := models.Users{}

		for rows.Next() {
			user := models.User{}
			if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.HashPassword, &user.ImageURL, &user.CreatedOn); err != nil {
				return err // Exit if we get an error
			}

			// Append Employee to Employees
			result.Users = append(result.Users, user)
		}
		// Return Employees in JSON format
		return c.JSON(result)
}

func Logout(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func PostQuestion(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func PostAnswer(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func PostStory(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func GetQuestions(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func GetAnswers(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func GetStories(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func GetAQuestion(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func GetAnAnswer(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func GetAStory(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func GetBlogs(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func GetABlog(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

func PostBlog(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}