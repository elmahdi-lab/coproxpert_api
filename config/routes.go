package config

import (
	"github.com/gofiber/fiber/v2"
	"ithumans.com/coproxpert/controllers"
	"ithumans.com/coproxpert/middleware"
	"ithumans.com/coproxpert/models"
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

	// Check subscription to limit the number of unit groups the user can create.
	secureApi.Post("/:organizationID/unit-group", middleware.CheckSubscriptionLimit(models.UnitGroupLimit), controllers.CreateUnitGroupAction)
	secureApi.Get("/:organizationID/unit-group/:id", controllers.GetUnitGroupAction)
	secureApi.Put("/:organizationID/unit-group/:id", controllers.UpdateUnitGroupAction)
	secureApi.Delete("/:organizationID/unit-group/:id", controllers.DeleteUnitGroupAction)

	// Check subscription to limit the number of units the user can create.
	secureApi.Post("/:unitGroupID/unit", middleware.CheckSubscriptionLimit(models.UnitLimit), controllers.CreateUnitAction)
	secureApi.Get("/:unitGroupID/unit/:id", controllers.GetUnitAction)
	secureApi.Put("/:unitGroupID/unit/:id", controllers.UpdateUnitAction)
	secureApi.Delete("/:unitGroupID/unit/:id", controllers.DeleteUnitAction)

	secureApi.Post("/:organizationID/maintenance", controllers.CreateMaintenanceAction)
	secureApi.Get("/:organizationID/maintenance/:id", controllers.GetMaintenanceAction)
	secureApi.Put("/:organizationID/maintenance/:id", controllers.UpdateMaintenanceAction)
	secureApi.Delete("/:organizationID/maintenance/:id", controllers.DeleteMaintenanceAction)

	secureApi.Post("/:organizationID/resolution", controllers.CreateResolutionAction)
	secureApi.Get("/:organizationID/resolution/:id", controllers.GetResolutionAction)
	secureApi.Put("/:organizationID/resolution/:id", controllers.UpdateResolutionAction)
	secureApi.Delete("/resolution/:id", controllers.DeleteResolutionAction)

	secureApi.Post("/:resolutionID/vote", controllers.CreateVoteAction)
	secureApi.Get("/:resolutionID/vote/:id", controllers.GetVoteAction)
	secureApi.Put("/:resolutionID/vote/:id", controllers.UpdateVoteAction)
	secureApi.Delete("/:resolutionID/vote/:id", controllers.DeleteVoteAction)

	secureApi.Post("/:organizationID/complaint", controllers.CreateComplaintAction)
	secureApi.Get("/:organizationID/complaint/:id", controllers.GetComplaintAction)
	secureApi.Put("/:organizationID/complaint/:id", controllers.UpdateComplaintAction)
	secureApi.Delete("/:organizationID/complaint/:id", controllers.DeleteComplaintAction)

	secureApi.Post("/file", controllers.UploadFileAction)
	secureApi.Get("/file/:id", controllers.GetFileAction)
	secureApi.Put("/file/:id", controllers.UpdateFileAction)
	secureApi.Delete("/file/:id", controllers.DeleteFileAction)

	slog.Info("Routes registered")
}
