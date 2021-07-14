package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sidmohanty11/gradbook/server/models"
)

// POST route for posting a question's answer.
func PostAnswer(c *fiber.Ctx) error {
	a := new(models.Answer)

	if err := c.BodyParser(a); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if a.AnswerText == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "nothing received."})
	}

	_, err := Psql.Exec("INSERT INTO answers (q_id, user_id, a_text) VALUES ($1, $2, $3)", a.QId, a.UserId, a.AnswerText)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Error posting answer to db."})
	}

	return c.JSON(a)
}

// GET route for getting a question's answers by the specific question_id.
func GetAnswers(c *fiber.Ctx) error {
	paramId := c.Params("id")
	rows, err := Psql.Query("SELECT answers.id, answers.user_id, answers.q_id, answers.a_text, users.username FROM answers LEFT JOIN users ON user_id = users.id WHERE q_id = $1;", paramId)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := models.Answers{}

	for rows.Next() {
		a := models.Answer{}

		if err := rows.Scan(&a.ID, &a.UserId, &a.QId, &a.AnswerText, &a.Username); err != nil {
			return err
		}

		result.Answers = append(result.Answers, a)
	}

	return c.JSON(result)
}
