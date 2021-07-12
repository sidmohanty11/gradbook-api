package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sidmohanty11/gradbook/server/models"
)

// POST route to post a question.
func PostQuestion(c *fiber.Ctx) error {
	q := new(models.Question)

	if err := c.BodyParser(q); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := Psql.Exec("INSERT INTO questions (q_text, user_id) VALUES ($1, $2)", q.Question, q.UserId)

	if err != nil {
		return err
	}

	return c.JSON(q)
}

// GET route to get all the questions, particularly fetching questions at the home page.
func GetQuestions(c *fiber.Ctx) error {
	rows, err := Psql.Query("SELECT questions.id, questions.user_id, questions.q_text, questions.created_on, users.username, users.image_url FROM questions LEFT JOIN users ON user_id = users.id ORDER BY created_on DESC;")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := models.Questions{}

	for rows.Next() {
		q := models.Question{}

		if err := rows.Scan(&q.ID, &q.UserId, &q.Question, &q.CreatedOn, &q.Username, &q.ImageURL); err != nil {
			return err
		}
		result.Questions = append(result.Questions, q)
	}

	return c.JSON(result)
}

// UPDATE route for changing question text.
func PutQuestion(c *fiber.Ctx) error {
	q := new(models.Question)

	if err := c.BodyParser(q); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := Psql.Query("UPDATE questions SET q_text=$1 WHERE id=$2", q.Question, q.ID)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(q)
}

// DELETE route for deleting the question from db and all answers related to it.
func DeleteQuestion(c *fiber.Ctx) error {
	paramId := c.Params("id")
	q := new(models.Question)

	if err := c.BodyParser(q); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := Psql.Exec("DELETE FROM questions WHERE id = $1", paramId)
	if err != nil {
		return err
	}

	_, err = Psql.Exec("DELETE FROM answers WHERE q_id = $1", paramId)
	if err != nil {
		return err
	}

	return c.JSON("Deleted")
}
