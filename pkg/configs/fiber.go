package configs

import (
	"github.com/helply/backend/platform/server"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig() fiber.Config {
	// Define server settings.
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	// Return Fiber configuration.
	return fiber.Config{
		ReadTimeout:  time.Second * time.Duration(readTimeoutSecondsCount),
		Prefork:      false,
		ErrorHandler: server.ErrorHandler,
	}
}
