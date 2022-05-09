package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/helply/backend/app/controllers"
)

func HelloRoutes(app *fiber.App) {
	app.Get("/", controllers.Hello)
}
