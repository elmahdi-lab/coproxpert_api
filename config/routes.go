package config

import (
	"github.com/gofiber/fiber/v2"
	"ithumans.com/coproxpert/controllers"
	"ithumans.com/coproxpert/middleware"
	"ithumans.com/coproxpert/repositories"
	"log/slog"
)

func RegisterRoutes(app *fiber.App) {

	app.Get("/healthcheck", controllers.HealthCheck)
	app.Post("/api/user/login", controllers.LoginAction)
	app.Post("/api/user/register", controllers.CreateUserAction)
	app.Post("/api/user/password-forget", controllers.PasswordForgetAction)
	app.Post("/api/user/password-reset", controllers.PasswordResetAction)

	// Group with authentication middleware for secure routes
	secureApi := app.Group("/api", middleware.AuthMiddleware)

	secureApi.Post("/user", controllers.CreateUserAction)
	secureApi.Get("/user/:id", middleware.IsOwnerMiddleware(repositories.UserType), controllers.GetUserAction)
	secureApi.Put("/user/:id", controllers.UpdateUserAction)
	secureApi.Delete("/user/:id", middleware.IsOwnerMiddleware(repositories.UserType), controllers.DeleteUserAction)
	secureApi.Put("/user/password/:id", middleware.IsOwnerMiddleware(repositories.UserType), controllers.UpdatePasswordAction)
	secureApi.Post("/user/logout", controllers.LogoutAction)

	secureApi.Post("/organization", controllers.CreateOrganizationAction)
	secureApi.Get("/organization/:id", controllers.GetOrganizationAction)
	secureApi.Put("/organization/:id", controllers.UpdateOrganizationAction)
	secureApi.Delete("/organization/:id", controllers.DeleteOrganizationAction)

	secureApi.Post("/unit-group", controllers.CreateUnitGroupAction)
	secureApi.Get("/unit-group/:id", controllers.GetUnitGroupAction)
	secureApi.Put("/unit-group/:id", controllers.UpdateUnitGroupAction)
	secureApi.Delete("/unit-group/:id", controllers.DeleteUnitGroupAction)

	secureApi.Post("/unit", controllers.CreateUnitAction)
	secureApi.Get("/unit/:id", controllers.GetUnitAction)
	secureApi.Put("/unit/:id", controllers.UpdateUnitAction)
	secureApi.Delete("/unit/:id", controllers.DeleteUnitAction)

	secureApi.Post("/maintenance", controllers.CreateMaintenanceAction)
	secureApi.Get("/maintenance/:id", controllers.GetMaintenanceAction)
	secureApi.Put("/maintenance/:id", controllers.UpdateMaintenanceAction)
	secureApi.Delete("/maintenance/:id", controllers.DeleteMaintenanceAction)

	secureApi.Post("/resolution", controllers.CreateResolutionAction)
	secureApi.Get("/resolution/:id", controllers.GetResolutionAction)
	secureApi.Put("/resolution/:id", controllers.UpdateResolutionAction)
	secureApi.Delete("/resolution/:id", controllers.DeleteResolutionAction)

	secureApi.Post("/vote", controllers.CreateVoteAction)
	secureApi.Get("/vote/:id", controllers.GetVoteAction)
	secureApi.Put("/vote/:id", controllers.UpdateVoteAction)
	secureApi.Delete("/vote/:id", controllers.DeleteVoteAction)

	secureApi.Post("/complaint", controllers.CreateComplaintAction)
	secureApi.Get("/complaint/:id", controllers.GetComplaintAction)
	secureApi.Put("/complaint/:id", controllers.UpdateComplaintAction)
	secureApi.Delete("/complaint/:id", controllers.DeleteComplaintAction)

	secureApi.Post("/file", controllers.UploadFileAction)
	secureApi.Get("/file/:id", controllers.GetFileAction)
	secureApi.Put("/file/:id", controllers.UpdateFileAction)
	secureApi.Delete("/file/:id", controllers.DeleteFileAction)

	slog.Info("Routes registered")
}
