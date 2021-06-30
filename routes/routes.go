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

	//post routes
	app.Post("/api/v1/question", handlers.PostQuestion)
	app.Post("/api/v1/answer", handlers.PostAnswer)
	app.Post("/api/v1/story", handlers.PostStory)

	//get routes
	app.Get("/api/v1/question", handlers.GetQuestion)
	app.Get("/api/v1/answer", handlers.GetAnswer)
	app.Get("/api/v1/story", handlers.GetStory)

	//get routes (specific)
	app.Get("/api/v1/question/:id", handlers.GetAQuestion)
	app.Get("/api/v1/answer/:id", handlers.GetAnAnswer)
	app.Get("/api/v1/story/:id", handlers.GetAStory)
}
