package config

import (
	"github.com/gofiber/fiber/v2"
	"ithumans.com/coproxpert/controllers"
	"ithumans.com/coproxpert/middleware/security"
)

func RegisterRoutes(app *fiber.App) {
	publicApi := app.Group("/public")
	publicApi.Get("/login", controllers.Login)
	publicApi.Get("/logout", controllers.Logout)
	publicApi.Get("/register", controllers.Register)
	publicApi.Get("/password-forget", controllers.PasswordForget)
	publicApi.Get("/healthcheck", controllers.HealthCheck)

	// Group with authentication middleware for secure routes
	secureApi := app.Group("/api", security.AuthMiddleware)
	secureApi.Get("/", controllers.UserGreeting)
}
