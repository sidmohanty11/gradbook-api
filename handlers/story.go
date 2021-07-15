package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sidmohanty11/gradbook/server/models"
)

func PostStory(c *fiber.Ctx) error {
	s := new(models.Story)

	_, err := ValidToken(c)

	if err != nil {
		return c.JSON(fiber.Map{"status": "unauthorized"})
	}

	if err := c.BodyParser(s); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if s.Branch == "" || s.Clubs == "" || s.ImageURL == "" || s.Motto == "" || s.Name == "" || s.Journey == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "fill the essential stuffs before submitting."})
	}

	_, err = Psql.Exec("INSERT INTO stories (user_id, name, branch, clubs, image_url, motto, github_link, linkedin_link, youtube_link, journey) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", s.UserId, s.Name, s.Branch, s.Clubs, s.ImageURL, s.Motto, s.GithubLink, s.LinkedinLink, s.YoutubeLink, s.Journey)

	if err != nil {
		return err
	}

	return c.JSON(s)
}

func GetStories(c *fiber.Ctx) error {
	_, err := ValidToken(c)

	if err != nil {
		return c.JSON(fiber.Map{"status": "unauthorized"})
	}

	rows, err := Psql.Query("SELECT stories.id, stories.user_id, stories.name, stories.branch, stories.clubs, stories.image_url, stories.motto, stories.github_link, stories.linkedin_link, stories.youtube_link, stories.journey, users.username, users.image_url FROM stories LEFT JOIN users ON user_id = users.id ORDER BY stories.id DESC;")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := models.Stories{}

	for rows.Next() {
		story := models.Story{}
		if err := rows.Scan(&story.ID, &story.UserId, &story.Name, &story.Branch, &story.Clubs, &story.ImageURL, &story.Motto, &story.GithubLink, &story.LinkedinLink, &story.YoutubeLink, &story.Journey, &story.Username, &story.UserImageUrl); err != nil {
			return err
		}
		result.Stories = append(result.Stories, story)
	}
	return c.JSON(result)
}

func GetAStory(c *fiber.Ctx) error {
	_, err := ValidToken(c)

	if err != nil {
		return c.JSON(fiber.Map{"status": "unauthorized"})
	}

	thestory := c.Params("id")

	row := Psql.QueryRow("SELECT stories.id, stories.user_id, stories.name, stories.branch, stories.clubs, stories.image_url, stories.motto, stories.github_link, stories.linkedin_link, stories.youtube_link, stories.journey, users.username, users.image_url FROM stories LEFT JOIN users ON user_id = users.id WHERE stories.id = $1;", thestory)

	story := models.Story{}
	if err := row.Scan(&story.ID, &story.UserId, &story.Name, &story.Branch, &story.Clubs, &story.ImageURL, &story.Motto, &story.GithubLink, &story.LinkedinLink, &story.YoutubeLink, &story.Journey, &story.Username, &story.UserImageUrl); err != nil {
		return err
	}

	return c.JSON(story)
}

func PutStory(c *fiber.Ctx) error {
	_, err := ValidToken(c)

	if err != nil {
		return c.JSON(fiber.Map{"status": "unauthorized"})
	}

	s := new(models.Story)

	if err := c.BodyParser(s); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if s.Branch == "" || s.Clubs == "" || s.ImageURL == "" || s.Motto == "" || s.Name == "" || s.Journey == "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "fill the essential stuffs before submitting."})
	}

	_, err = Psql.Query("UPDATE stories SET name=$1, branch=$2, clubs=$3, image_url=$4, motto=$5, github_link=$6, linkedin_link=$7, youtube_link=$8, journey=$9 WHERE id=$10", s.Name, s.Branch, s.Clubs, s.ImageURL, s.Motto, s.GithubLink, s.LinkedinLink, s.YoutubeLink, s.Journey, s.ID)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(s)
}
