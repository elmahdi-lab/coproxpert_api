package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func RegisterRoutes(app *fiber.App) {
	// Define your routes here
	app.Get("/", func(c *fiber.Ctx) error {
		log.Info("Hello world!")
		return c.JSON(fiber.Map{"message": "pong"})
	})
}
