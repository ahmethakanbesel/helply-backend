package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/helply/backend/app/controllers"
	"github.com/helply/backend/pkg/middleware"
)

// ApiRoutes func for describe group of public routes.
func ApiRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for auth
	route.Post("/auth/login", controllers.Login)
	// Example restricted route
	route.Get("/restricted", middleware.Protected(), controllers.Hello)

	// Post routes
	route.Get("/posts", controllers.GetPosts)
	route.Get("/posts/:id", controllers.GetPost)
	route.Post("/posts/", controllers.CreatePost)
	route.Post("/posts/:id", controllers.UpdatePost)
	route.Delete("/posts/:id", controllers.DeletePost)

	route.Get("/articles", controllers.GetArticles)
	route.Get("/articles/:id", controllers.GetArticle)
	route.Post("/articles/", controllers.CreateArticle)
	route.Post("/articles/:id", controllers.UpdateArticle)
	route.Delete("/articles/:id", controllers.DeleteArticle)

}
