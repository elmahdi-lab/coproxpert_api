package config

import (
	"github.com/gofiber/fiber/v2"
	"ithumans.com/coproxpert/controllers"
	"ithumans.com/coproxpert/middleware"
)

func RegisterRoutes(app *fiber.App) {

	app.Get("/healthcheck", controllers.HealthCheck)
	app.Post("/api/user/login", controllers.LoginAction)
	app.Post("/api/user/logout", controllers.LogoutAction)
	app.Post("/api/user/register", controllers.CreateUserAction)
	app.Post("/api/user/password-forget", controllers.PasswordForgetAction)
	app.Post("/api/user/password-reset", controllers.PasswordResetAction)

	// Group with authentication middleware for secure routes
	secureApi := app.Group("/api", middleware.AuthMiddleware)

	secureApi.Post("/user", controllers.CreateUserAction)
	secureApi.Get("/user/:id", controllers.GetUserAction)
	secureApi.Put("/user", controllers.UpdateUserAction)
	secureApi.Delete("/user/:id", controllers.DeleteUserAction)

	secureApi.Post("/contact", controllers.CreateContactAction)
	secureApi.Get("/contact/:id", controllers.GetContactAction)
	secureApi.Put("/contact", controllers.UpdateContactAction)
	secureApi.Delete("/contact/:id", controllers.DeleteContactAction)

	secureApi.Post("/organization", controllers.CreateOrganizationAction)
	secureApi.Get("/organization/:id", controllers.GetOrganizationAction)
	secureApi.Put("/organization", controllers.UpdateOrganizationAction)
	secureApi.Delete("/organization/:id", controllers.DeleteOrganizationAction)

	secureApi.Post("/unit-group", controllers.CreateUnitGroupAction)
	secureApi.Get("/unit-group/:id", controllers.GetUnitGroupAction)
	secureApi.Put("/unit-group", controllers.UpdateUnitGroupAction)
	secureApi.Delete("/unit-group/:id", controllers.DeleteUnitGroupAction)

	secureApi.Post("/unit", controllers.CreateUnitAction)
	secureApi.Get("/unit/:id", controllers.GetUnitAction)
	secureApi.Put("/unit", controllers.UpdateUnitAction)
	secureApi.Delete("/unit/:id", controllers.DeleteUnitAction)

	secureApi.Post("/maintenance", controllers.CreateMaintenanceAction)
	secureApi.Get("/maintenance/:id", controllers.GetMaintenanceAction)
	secureApi.Put("/maintenance", controllers.UpdateMaintenanceAction)
	secureApi.Delete("/maintenance/:id", controllers.DeleteMaintenanceAction)

	secureApi.Post("/resolution", controllers.CreateResolutionAction)
	secureApi.Get("/resolution/:id", controllers.GetResolutionAction)
	secureApi.Put("/resolution", controllers.UpdateResolutionAction)
	secureApi.Delete("/resolution/:id", controllers.DeleteResolutionAction)

	secureApi.Post("/vote", controllers.CreateVoteAction)
	secureApi.Get("/vote/:id", controllers.GetVoteAction)
	secureApi.Put("/vote", controllers.UpdateVoteAction)
	secureApi.Delete("/vote/:id", controllers.DeleteVoteAction)

}
