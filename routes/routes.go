package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sidmohanty11/gradbook/server/handlers"
)

func Setup(app *fiber.App) {
	//auth routes
	app.Post("/api/v1/register", handlers.Register)
	app.Post("/api/v1/login", handlers.Login)
	app.Post("/api/v1/logout", handlers.Logout)
	app.Get("/api/v1/user", handlers.User)

	//story routes
	app.Get("/api/v1/story", handlers.GetStories)
	app.Get("/api/v1/story/:id", handlers.GetAStory)
	app.Post("/api/v1/story", handlers.PostStory)
	
	//blog routes
	app.Get("/api/v1/story", handlers.GetBlogs)
	app.Get("/api/v1/story/:id", handlers.GetABlog)
	app.Post("/api/v1/story", handlers.PostBlog)
	
	//qa routes
	app.Get("/api/v1/questions", handlers.GetQuestions)
	app.Get("/api/v1/answers", handlers.GetAnswers)
	app.Post("/api/v1/question", handlers.PostQuestion)
	app.Get("/api/v1/question/:id", handlers.GetAQuestion)
	app.Post("/api/v1/answer", handlers.PostAnswer)
	app.Get("/api/v1/answer/:id", handlers.GetAnAnswer)
}
