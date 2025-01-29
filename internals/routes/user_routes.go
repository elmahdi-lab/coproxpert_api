package routes

import (
	"github.com/gofiber/fiber/v2"
	"ithumans.com/coproxpert/data/controllers"
)

func RegisterUserRoutes(app *fiber.App) {
	app.Post("/api/user/login", controllers.LoginAction)
	app.Post("/api/user/register", controllers.CreateUserAction)
	app.Post("/api/user/password-forget", controllers.PasswordForgetAction)
	app.Post("/api/user/password-reset", controllers.PasswordResetAction)
	app.Get("/api/user/refresh-token/:id", controllers.RefreshJWTTokenAction)
	// api := app.Group("/api/user", middleware.AuthMiddleware)

	//api.Get("/profile", controllers.GetProfileAction)

}
