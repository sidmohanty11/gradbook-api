package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sidmohanty11/gradbook/server/handlers"
	"github.com/sidmohanty11/gradbook/server/middleware"
)

func Setup(app *fiber.App) {
	//auth routes
	app.Post("/api/v1/register", handlers.Register)
	app.Post("/api/v1/login", handlers.Login)
	app.Get("/api/v1/users/:id", handlers.UserById)
	app.Get("/api/v1/user/:username", handlers.UserByUsername)
	app.Delete("/api/v1/user/:id", handlers.DeleteUser)

	//story routes
	app.Get("/api/v1/stories", handlers.GetStories)
	app.Get("/api/v1/story/:id", handlers.GetAStory)
	app.Post("/api/v1/story", middleware.Protected(), handlers.PostStory)
	app.Put("/api/v1/story/:id", middleware.Protected(), handlers.PutStory)

	//blog routes
	app.Get("/api/v1/blogs", handlers.GetBlogs)
	app.Get("/api/v1/blogs/:id", handlers.GetBlogsByUserId)
	app.Get("/api/v1/blog/:id", handlers.GetABlog)
	app.Post("/api/v1/blog", middleware.Protected(), handlers.PostBlog)
	app.Put("/api/v1/blog/:id", middleware.Protected(), handlers.PutBlog)
	app.Delete("/api/v1/blog/:id", middleware.Protected(), handlers.DeleteBlog)

	//qa routes
	app.Get("/api/v1/questions", handlers.GetQuestions)
	app.Get("/api/v1/answers/:id", handlers.GetAnswers)
	app.Post("/api/v1/question", middleware.Protected(), handlers.PostQuestion)
	app.Post("/api/v1/answer", middleware.Protected(), handlers.PostAnswer)
	app.Put("/api/v1/question/:id", middleware.Protected(), handlers.PutQuestion)
	app.Delete("/api/v1/question/:id", middleware.Protected(), handlers.DeleteQuestion)

	//messages route
	app.Post("/api/v1/messages/:from/:to", middleware.Protected(), handlers.PostMessages)
	app.Get("/api/v1/messages/:id/:toid", handlers.GetMessages)
}
