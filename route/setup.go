package route

import (
	"credit-plus/internal/exception"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"strings"
)

func SetupApp() *fiber.App {
	app := fiber.New(fiber.Config{
		BodyLimit:     3 * 1024 * 1024,
		CaseSensitive: true,
		StrictRouting: true,
		ErrorHandler:  exception.NewHTTPErrorHandler,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodPut,
		}, ","),
	}))
	app.Use(logger.New())
	app.Use(recover.New())

	return app
}
