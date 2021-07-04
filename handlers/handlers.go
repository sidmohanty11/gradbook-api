package handlers

import (
	"database/sql"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/sidmohanty11/gradbook/server/db"
	"github.com/sidmohanty11/gradbook/server/helpers"
	"github.com/sidmohanty11/gradbook/server/models"
	"golang.org/x/crypto/bcrypt"
)

var psql *sql.DB

//creates a new repository
func NewRepo(db *db.DB) {
	psql = db.SQL
}

// ------------------AUTH HANDLERS----------------------------------
func Register(c *fiber.Ctx) error {
	type RegisterInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
		ImageURL string `json:"image_url"`
	}
	var input RegisterInput
	if err := c.BodyParser(&input); err != nil {
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

	_, err = psql.Exec("insert into users (username, image_url, email) values ($1,$2,$3);", identity, imgurl, email)

	if err != nil {
		return err
	}

	_, err = psql.Exec("insert into login (username, hash) values ($1,$2);", identity, hashedpass)

	if err != nil {
		return err
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = identity
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}

func Login(c *fiber.Ctx) error {
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

	row := psql.QueryRow("SELECT username, hash FROM login WHERE username = $1;", identity)

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

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}

// User HANDLER ----------------------------------------------------------------
func User(c *fiber.Ctx) error {
	rows, err := psql.Query("SELECT id, username, email, image_url, created_on, last_login FROM users;")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := models.Users{}

	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.ImageURL, &user.CreatedOn, &user.LastLogin); err != nil {
			return err
		}
		result.Users = append(result.Users, user)
	}
	return c.JSON(result)
}

func DeleteUser(c *fiber.Ctx) error {
	paramId := c.Params("id")
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := psql.Exec("DELETE FROM users WHERE id = $1", paramId)
	if err != nil {
		return err
	}

	_, err = psql.Exec("DELETE FROM login WHERE user_id = $1", paramId)
	if err != nil {
		return err
	}

	return c.JSON("Deleted")
}

// Question/Answer HANDLERS----------------------------------------------------------
func PostQuestion(c *fiber.Ctx) error {
	q := new(models.Question)

	if err := c.BodyParser(q); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := psql.Exec("INSERT INTO questions (q_text, user_id) VALUES ($1, $2)", q.Question, q.UserId)

	if err != nil {
		return err
	}

	return c.JSON(q)
}

func PostAnswer(c *fiber.Ctx) error {
	a := new(models.Answer)

	if err := c.BodyParser(a); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := psql.Exec("INSERT INTO answers (q_id, user_id, a_text) VALUES ($1, $2, $3)", a.QId, a.UserId, a.AnswerText)

	if err != nil {
		return err
	}

	return c.JSON(a)
}

func GetQuestions(c *fiber.Ctx) error {
	rows, err := psql.Query("select id, user_id, q_text, created_on from questions;")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := models.Questions{}

	for rows.Next() {
		q := models.Question{}

		if err := rows.Scan(&q.ID, &q.UserId, &q.Question, &q.CreatedOn); err != nil {
			return err
		}
		result.Questions = append(result.Questions, q)
	}

	return c.JSON(result)
}

func GetAnswers(c *fiber.Ctx) error {
	rows, err := psql.Query("select id, user_id, q_id, a_text from answers;")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := models.Answers{}

	for rows.Next() {
		a := models.Answer{}

		if err := rows.Scan(&a.ID, &a.UserId, &a.QId, &a.AnswerText); err != nil {
			return err
		}

		result.Answers = append(result.Answers, a)
	}

	return c.JSON(result)
}

func PutQuestion(c *fiber.Ctx) error {
	q := new(models.Question)

	if err := c.BodyParser(q); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := psql.Query("UPDATE questions SET q_text=$1 WHERE id=$2", q.Question, q.ID)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(q)
}

func DeleteQuestion(c *fiber.Ctx) error {
	paramId := c.Params("id")
	q := new(models.Question)

	if err := c.BodyParser(q); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := psql.Exec("DELETE FROM questions WHERE id = $1", paramId)
	if err != nil {
		return err
	}

	_, err = psql.Exec("DELETE FROM answers WHERE q_id = $1", paramId)
	if err != nil {
		return err
	}

	return c.JSON("Deleted")
}

// Story HANDLERS----------------------------------------------------------------
func PostStory(c *fiber.Ctx) error {
	s := new(models.Story)

	if err := c.BodyParser(s); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := psql.Exec("INSERT INTO stories (user_id, name, branch, clubs, image_url, motto, github_link, linkedin_link, youtube_link, journey) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", s.UserId, s.Name, s.Branch, s.Clubs, s.ImageURL, s.Motto, s.GithubLink, s.LinkedinLink, s.YoutubeLink, s.Journey)

	if err != nil {
		return err
	}

	return c.JSON(s)
}

func GetStories(c *fiber.Ctx) error {
	rows, err := psql.Query("SELECT id, user_id, name, branch, clubs, image_url, motto, github_link, linkedin_link, youtube_link, journey FROM stories;")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := models.Stories{}

	for rows.Next() {
		story := models.Story{}
		if err := rows.Scan(&story.ID, &story.UserId, &story.Name, &story.Branch, &story.Clubs, &story.ImageURL, &story.Motto, &story.GithubLink, &story.LinkedinLink, &story.YoutubeLink, &story.Journey); err != nil {
			return err
		}
		result.Stories = append(result.Stories, story)
	}
	return c.JSON(result)
}

func GetAStory(c *fiber.Ctx) error {
	thestory := c.Params("id")

	row := psql.QueryRow("SELECT id, user_id, name, branch, clubs, image_url, motto, github_link, linkedin_link, youtube_link, journey FROM stories WHERE id = $1;", thestory)

	story := models.Story{}
	if err := row.Scan(&story.ID, &story.UserId, &story.Name, &story.Branch, &story.Clubs, &story.ImageURL, &story.Motto, &story.GithubLink, &story.LinkedinLink, &story.YoutubeLink, &story.Journey); err != nil {
		return err
	}

	return c.JSON(story)
}

func PutStory(c *fiber.Ctx) error {
	s := new(models.Story)

	if err := c.BodyParser(s); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := psql.Query("UPDATE stories SET name=$1, branch=$2, clubs=$3, image_url=$4, motto=$5, github_link=$6, linkedin_link=$7, youtube_link=$8, journey=$9 WHERE id=$10",s.Name,s.Branch,s.Clubs,s.ImageURL,s.Motto,s.GithubLink,s.LinkedinLink,s.YoutubeLink,s.Journey, s.ID)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(s)
}

// Blog HANDLERS-------------------------------------------------------------
func GetBlogs(c *fiber.Ctx) error {
	rows, err := psql.Query("SELECT id, user_id, blog_title, blog_text, created_on FROM blogs;")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := models.Blogs{}

	for rows.Next() {
		blog := models.Blog{}
		if err := rows.Scan(&blog.ID, &blog.UserId, &blog.BlogTitle, &blog.BlogText, &blog.CreatedOn); err != nil {
			return err
		}
		result.Blogs = append(result.Blogs, blog)
	}
	return c.JSON(result)
}

func GetABlog(c *fiber.Ctx) error {
	theBlog := c.Params("id")

	row := psql.QueryRow("SELECT id, user_id, blog_title, blog_text, created_on FROM blogs WHERE id = $1;", theBlog)

	blog := models.Blog{}
	if err := row.Scan(&blog.ID, &blog.UserId, &blog.BlogTitle, &blog.BlogText, &blog.CreatedOn); err != nil {
		return err
	}

	return c.JSON(blog)
}

func PostBlog(c *fiber.Ctx) error {
	b := new(models.Blog)

	if err := c.BodyParser(b); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := psql.Exec("INSERT INTO stories (user_id, blog_title, blog_text) VALUES ($1, $2, $3)", b.UserId, b.BlogTitle, b.BlogText)

	if err != nil {
		return err
	}

	return c.JSON(b)
}

func PutBlog(c *fiber.Ctx) error {
	b := new(models.Blog)

	if err := c.BodyParser(b); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := psql.Query("UPDATE blogs SET blog_title=$1, blog_text=$2 WHERE id=$3", b.BlogTitle,b.BlogText,b.ID)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(b)
}

func DeleteBlog(c *fiber.Ctx) error {
	paramId := c.Params("id")
	b := new(models.Blog)

	if err := c.BodyParser(b); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := psql.Exec("DELETE FROM blogs WHERE id = $1", paramId)
	if err != nil {
		return err
	}

	return c.JSON("Deleted")
}
