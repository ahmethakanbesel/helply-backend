package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/helply/backend/pkg/configs"
	"github.com/helply/backend/pkg/middleware"
	"github.com/helply/backend/pkg/routes"
	"log"
)

func CreateServer(port int) {
	// Create Fiber App
	config := configs.FiberConfig()
	app := fiber.New(config)

	// Middlewares
	//app.Use(middleware.Example)
	middleware.FiberMiddleware(app)

	// Mount routes
	routes.HelloRoutes(app)
	routes.ApiRoutes(app)
	routes.SwaggerRoute(app)
	routes.DashboardRoute(app)
	routes.NotFoundRoute(app)

	// Start server
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
