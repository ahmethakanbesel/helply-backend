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
	route.Post("/articles/save", controllers.CreateUserSavedArticle)
	route.Delete("/articles/:id", controllers.DeleteArticle)

	// Article vote routes
	route.Post("/articles/:id/vote", controllers.VoteArticle)

	// Article Tag routes
	route.Get("/article-tags", controllers.GetArticleTags)
	route.Get("/article-tags/:id", controllers.GetArticleTag)
	route.Post("/article-tags/", controllers.CreateArticleTag)
	route.Post("/article-tags/:id", controllers.UpdateArticleTag)
	route.Delete("/article-tags/:id", controllers.DeleteArticleTag)

	// Article Category routes
	route.Get("/article-categories", controllers.GetArticleCategories)
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

	// License routes
	route.Get("/licences", controllers.GetLicenses)
	route.Get("/licences/:id", controllers.GetLicense)
	route.Post("/licences/", controllers.CreateLicense)
	route.Post("/licences/:id", controllers.UpdateLicense)
	route.Delete("/licences/:id", controllers.DeleteLicense)

	// Customer License routes
	route.Get("/customer-licenses", controllers.GetCustomerLicenses)
	route.Get("/customer-licenses/:id", controllers.GetCustomerLicense)
	route.Post("/customer-licenses/", controllers.CreateCustomerLicense)
	route.Post("/customer-licenses/:id", controllers.UpdateCustomerLicense)
	route.Delete("/customer-licenses/:id", controllers.DeleteCustomerLicense)

	// Ticket routes
	route.Get("/tickets", controllers.GetTickets)
	route.Get("/tickets/:id", controllers.GetTicket)
	route.Post("/tickets/", controllers.CreateTicket)
	route.Post("/tickets/:id", controllers.UpdateTicket)
	route.Post("/tickets/:id/close", controllers.CloseTicket)
	route.Delete("/tickets/:id", controllers.DeleteTicket)

	// Ticket Reply routes
	route.Get("/ticket-replies/", controllers.GetTicketReplies)
	route.Get("/ticket-replies/:id", controllers.GetTicketRepliesById)
	route.Post("/ticket-replies/", controllers.CreateTicketReply)

	// Ticket Topic routes
	route.Get("/ticket-topics/", controllers.GetTicketTopics)

	// Stat routes
	route.Get("/admin-stats", controllers.GetAdminStats)

	// File Routes
	route.Post("/files/", controllers.UploadFile)
}
