package config

import (
	"github.com/gofiber/fiber/v2"
	"ithumans.com/coproxpert/controllers"
	"ithumans.com/coproxpert/middleware"
)

func RegisterRoutes(app *fiber.App) {

	app.Get("/healthcheck", controllers.HealthCheck)
	app.Post("/api/user/login", controllers.LoginAction)
	app.Post("/api/user/register", controllers.CreateUserAction)
	app.Post("/api/user/password-forget", controllers.PasswordForgetAction)

	// Group with authentication middleware for secure routes
	secureApi := app.Group("/api", middleware.AuthMiddleware)

	secureApi.Get("user/logout", controllers.LogoutAction)

	secureApi.Post("/user", controllers.CreateUserAction)
	secureApi.Get("/user/:id", controllers.GetUserAction)
	secureApi.Put("/user", controllers.UpdateUserAction)
	secureApi.Delete("/user/:id", controllers.DeleteUserAction)

	secureApi.Post("/contact", controllers.CreateContactAction)
	secureApi.Get("/contact/:id", controllers.GetContactAction)
	secureApi.Put("/contact", controllers.UpdateContactAction)
	secureApi.Delete("/contact/:id", controllers.DeleteContactAction)

}
