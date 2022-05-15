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

	// User routes
	route.Get("/users", controllers.GetUsers)
	route.Get("/users/:id", controllers.GetUser)
	route.Post("/users/", controllers.CreateUser)
	route.Post("/users/:id", controllers.UpdateUser)
	route.Delete("/users/:id", controllers.DeleteUser)
}
