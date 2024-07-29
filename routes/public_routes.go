package routes

import (
	"github.com/gofiber/fiber/v2"
	"ithumans.com/coproxpert/controllers"
)

func RegisterPublicRoutes(app *fiber.App) {
	app.Get("/healthcheck", controllers.HealthCheck)

}
