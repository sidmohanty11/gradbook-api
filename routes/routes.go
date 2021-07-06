package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sidmohanty11/gradbook/server/handlers"
)

func Setup(app *fiber.App) {
	//auth routes
	app.Post("/api/v1/register", handlers.Register)
	app.Post("/api/v1/login", handlers.Login)
	app.Get("/api/v1/users/:id", handlers.UserById)
	app.Get("/api/v1/user/:username", handlers.UserByUsername)
	app.Delete("/api/v1/story/:id", handlers.DeleteUser)

	//story routes
	app.Get("/api/v1/stories", handlers.GetStories)
	app.Get("/api/v1/story/:id", handlers.GetAStory)
	app.Post("/api/v1/story", handlers.PostStory)
	app.Put("/api/v1/story/:id", handlers.PutStory)
	
	//blog routes
	app.Get("/api/v1/blogs", handlers.GetBlogs)
	app.Get("/api/v1/blog/:id", handlers.GetABlog)
	app.Post("/api/v1/blog", handlers.PostBlog)
	app.Put("/api/v1/blog/:id", handlers.PutBlog)
	app.Delete("/api/v1/blog/:id", handlers.DeleteBlog)
	
	//qa routes
	app.Get("/api/v1/questions", handlers.GetQuestions)
	app.Get("/api/v1/answers/:id", handlers.GetAnswers)
	app.Post("/api/v1/question", handlers.PostQuestion)
	app.Post("/api/v1/answer", handlers.PostAnswer)
	app.Put("/api/v1/question/:id", handlers.PutQuestion)
	app.Delete("/api/v1/question/:id", handlers.DeleteQuestion)
}
