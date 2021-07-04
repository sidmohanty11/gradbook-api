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

func Register(c *fiber.Ctx) error {
	type RegisterInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email string `json:"email"`
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

	_, err = psql.Exec("insert into users (username, hash, image_url, email) values ($1,$2,$3,$4);", identity, hashedpass, imgurl, email)

	if err != nil {
		return err
	}

	_, err = psql.Query("insert into login (username, hash) values ($1,$2);", identity, hashedpass)

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

	rows, err := psql.Query("SELECT email, hash FROM login;")

	if err != nil {
		return err
	}

	var dbUsername string
	var dbPassword string

	if err := rows.Scan(&dbUsername, &dbPassword); err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(pass))

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

func GetQuestionsAndAnswers(c *fiber.Ctx) error {
	rows, err := psql.Query("select * from questions INNER JOIN answers on questions.id = answers.q_id;")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := models.Qas{}

	for rows.Next() {
		qa := models.Qa{}
		ans := models.Answer{}

		if err := rows.Scan(&qa.ID, &qa.UserId, &qa.Question, &qa.CreatedOn, &ans.ID, &ans.QId, &ans.UserId, &ans.AnswerText); err != nil {
			return err
		}
		
		qa.Answers = append(qa.Answers, ans)
		result.Qas = append(result.Qas, qa)
	}

	return c.JSON(result)
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
	return c.SendString("HELLO")
}
