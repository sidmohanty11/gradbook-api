package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sidmohanty11/gradbook/server/models"
)

func PostMessages(c *fiber.Ctx) error {
	fromUserID := c.Params("from")
	toUserID := c.Params("to")

	m := new(models.Message)

	if err := c.BodyParser(m); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if m.Content == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "fill the essential stuffs before submitting."})
	}

	_, err := Psql.Exec("INSERT INTO messages (from_user_id, to_user_id, content) VALUES ($1, $2, $3)", fromUserID, toUserID, m.Content)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(m)
}

func GetMessages(c *fiber.Ctx) error {
	userID := c.Params("id")
	toUserID := c.Params("toid")
	rows, err := Psql.Query("SELECT id, from_user_id, to_user_id, content, created_on FROM messages WHERE from_user_id = $1 AND to_user_id = $2;", userID, toUserID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := models.Messages{}

	for rows.Next() {
		message := models.Message{}
		if err := rows.Scan(&message.ID, &message.FromUserId, &message.ToUserId, &message.Content, &message.CreatedOn); err != nil {
			return err
		}
		result.Messages = append(result.Messages, message)
	}
	return c.JSON(result)
}
