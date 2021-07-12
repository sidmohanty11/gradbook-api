package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sidmohanty11/gradbook/server/models"
)

func GetBlogs(c *fiber.Ctx) error {
	rows, err := Psql.Query("SELECT blogs.id, blogs.user_id, blogs.blog_title, blogs.blog_text, blogs.created_on, users.username, users.image_url FROM blogs LEFT JOIN users ON user_id = users.id;")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := models.Blogs{}

	for rows.Next() {
		blog := models.Blog{}
		if err := rows.Scan(&blog.ID, &blog.UserId, &blog.BlogTitle, &blog.BlogText, &blog.CreatedOn, &blog.Username, &blog.ImageURL); err != nil {
			return err
		}
		result.Blogs = append(result.Blogs, blog)
	}
	return c.JSON(result)
}

func GetBlogsByUserId(c *fiber.Ctx) error {
	userId := c.Params("id")
	rows, err := Psql.Query("SELECT blogs.id, blogs.user_id, blogs.blog_title, blogs.blog_text, blogs.created_on, users.username, users.image_url FROM blogs LEFT JOIN users ON user_id = users.id WHERE user_id = $1;", userId)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := models.Blogs{}

	for rows.Next() {
		blog := models.Blog{}
		if err := rows.Scan(&blog.ID, &blog.UserId, &blog.BlogTitle, &blog.BlogText, &blog.CreatedOn, &blog.Username, &blog.ImageURL); err != nil {
			return err
		}
		result.Blogs = append(result.Blogs, blog)
	}
	return c.JSON(result)
}

func GetABlog(c *fiber.Ctx) error {
	theBlog := c.Params("id")

	row := Psql.QueryRow("SELECT id, user_id, blog_title, blog_text, created_on FROM blogs WHERE id = $1;", theBlog)

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

	_, err := Psql.Exec("INSERT INTO stories (user_id, blog_title, blog_text) VALUES ($1, $2, $3)", b.UserId, b.BlogTitle, b.BlogText)

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

	_, err := Psql.Query("UPDATE blogs SET blog_title=$1, blog_text=$2 WHERE id=$3", b.BlogTitle, b.BlogText, b.ID)
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

	_, err := Psql.Exec("DELETE FROM blogs WHERE id = $1", paramId)
	if err != nil {
		return err
	}

	return c.JSON("Deleted")
}
