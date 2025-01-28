package routes

import (
	"github.com/gofiber/fiber/v2"
	"ithumans.com/coproxpert/controllers"
	"ithumans.com/coproxpert/middleware"
	"ithumans.com/coproxpert/models"
)

func RegisterAdminRoutes(app *fiber.App) {

	// Group with authentication middleware for secure routes
	api := app.Group("/api", middleware.JWTProtected)

	api.Post("/user", controllers.CreateUserAction)
	api.Get("/user/:id", controllers.GetUserAction)
	api.Put("/user/:id", controllers.UpdateUserAction)
	api.Delete("/user/:id", controllers.DeleteUserAction)
	api.Put("/user/password/:id", controllers.UpdatePasswordAction)
	api.Post("/user/logout", controllers.LogoutAction)

	api.Post("/user/subscribe/:type", controllers.Subscribe)

	// Check subscription to limit the number of unit groups the user can create.
	api.Post("/unit-group", middleware.CheckSubscriptionLimit(models.UnitGroupLimit), controllers.CreateUnitGroupAction)
	api.Get("/unit-group/:id", middleware.ResourceAccess(
		models.ManagerRole, models.UnitGroupEntity), controllers.GetUnitGroupAction)
	api.Put("/unit-group/:id", middleware.ResourceAccess(
		models.ManagerRole, models.UnitGroupEntity), controllers.UpdateUnitGroupAction)
	api.Delete("/unit-group/:id", middleware.ResourceAccess(
		models.ManagerRole, models.UnitGroupEntity), controllers.DeleteUnitGroupAction)

	// Check subscription to limit the number of units the user can create.
	api.Post("/unit", middleware.CheckSubscriptionLimit(models.UnitLimit), controllers.CreateUnitAction)
	api.Get("/unit/:id", middleware.ResourceAccess(
		models.UserRole, models.UnitEntity), controllers.GetUnitAction)
	api.Put("/unit/:id", middleware.ResourceAccess(
		models.ManagerRole, models.UnitEntity), controllers.UpdateUnitAction)
	api.Delete("/unit/:id", middleware.ResourceAccess(
		models.ManagerRole, models.UnitEntity), controllers.DeleteUnitAction)

	api.Post("/maintenance", controllers.CreateMaintenanceAction)
	api.Get("/maintenance/:id", controllers.GetMaintenanceAction)
	api.Put("/maintenance/:id", controllers.UpdateMaintenanceAction)
	api.Delete("/maintenance/:id", controllers.DeleteMaintenanceAction)

	api.Post("/resolution", controllers.CreateResolutionAction)
	api.Get("/resolution/:id", controllers.GetResolutionAction)
	api.Put("/resolution/:id", controllers.UpdateResolutionAction)
	api.Delete("/resolution/:id", controllers.DeleteResolutionAction)

	api.Post("/vote", controllers.CreateVoteAction)
	api.Get("/vote/:id", controllers.GetVoteAction)
	api.Put("/vote/:id", controllers.UpdateVoteAction)
	api.Delete("/vote/:id", controllers.DeleteVoteAction)

	api.Post("/complaint", controllers.CreateComplaintAction)
	api.Get("/complaint/:id", controllers.GetComplaintAction)       // user or manager can view complaints.
	api.Put("/complaint/:id", controllers.UpdateComplaintAction)    // user or manager can update complaints.
	api.Delete("/complaint/:id", controllers.DeleteComplaintAction) // manager can delete

	api.Post("/file", controllers.UploadFileAction)
	api.Get("/file/:id", controllers.GetFileAction)
	api.Put("/file/:id", controllers.UpdateFileAction)
	api.Delete("/file/:id", controllers.DeleteFileAction)
}
