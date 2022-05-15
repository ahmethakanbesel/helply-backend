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

	// Article routes
	route.Get("/articles", controllers.GetArticles)
	route.Get("/articles/:id", controllers.GetArticle)
	route.Post("/articles/", controllers.CreateArticle)
	route.Post("/articles/:id", controllers.UpdateArticle)
	route.Delete("/articles/:id", controllers.DeleteArticle)

	// Article Tag routes
	route.Get("/article-tags", controllers.GetArticleTags)
	route.Get("/article-tags/:id", controllers.GetArticleTag)
	route.Post("/article-tags/", controllers.CreateArticleTag)
	route.Post("/article-tags/:id", controllers.UpdateArticleTag)
	route.Delete("/article-tags/:id", controllers.DeleteArticleTag)

	// Article Category routes
	route.Get("/article-categories", controllers.GetArticleCategorys)
	route.Get("/article-categories/:id", controllers.GetArticleCategory)
	route.Post("/article-categories/", controllers.CreateArticleCategory)
	route.Post("/article-categories/:id", controllers.UpdateArticleCategory)
	route.Delete("/article-categories/:id", controllers.DeleteArticleCategory)

	// Product routes
	route.Get("/products", controllers.GetProducts)
	route.Get("/products/:id", controllers.GetProduct)
	route.Post("/products/", controllers.CreateProduct)
	route.Post("/products/:id", controllers.UpdateProduct)
	route.Delete("/products/:id", controllers.DeleteProduct)
}
