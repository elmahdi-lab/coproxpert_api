package routes

import (
	"github.com/gofiber/fiber/v2"
	controllers2 "ithumans.com/coproxpert/data/controllers"
	models2 "ithumans.com/coproxpert/data/models"
	middleware2 "ithumans.com/coproxpert/internals/middleware"
)

func RegisterAdminRoutes(app *fiber.App) {

	// Group with authentication middleware for secure routes
	api := app.Group("/api", middleware2.JWTProtected)

	api.Post("/user", controllers2.CreateUserAction)
	api.Get("/user/:id", controllers2.GetUserAction)
	api.Put("/user/:id", controllers2.UpdateUserAction)
	api.Delete("/user/:id", controllers2.DeleteUserAction)
	api.Put("/user/password/:id", controllers2.UpdatePasswordAction)
	api.Post("/user/logout", controllers2.LogoutAction)

	api.Post("/user/subscribe/:type", controllers2.Subscribe)

	// Check subscription to limit the number of unit groups the user can create.
	api.Post("/unit-group", middleware2.CheckSubscriptionLimit(models2.UnitGroupLimit), controllers2.CreateUnitGroupAction)
	api.Get("/unit-group/:id", middleware2.ResourceAccess(
		models2.ManagerRole, models2.UnitGroupEntity), controllers2.GetUnitGroupAction)
	api.Put("/unit-group/:id", middleware2.ResourceAccess(
		models2.ManagerRole, models2.UnitGroupEntity), controllers2.UpdateUnitGroupAction)
	api.Delete("/unit-group/:id", middleware2.ResourceAccess(
		models2.ManagerRole, models2.UnitGroupEntity), controllers2.DeleteUnitGroupAction)

	// Check subscription to limit the number of units the user can create.
	api.Post("/unit", middleware2.CheckSubscriptionLimit(models2.UnitLimit), controllers2.CreateUnitAction)
	api.Get("/unit/:id", middleware2.ResourceAccess(
		models2.UserRole, models2.UnitEntity), controllers2.GetUnitAction)
	api.Put("/unit/:id", middleware2.ResourceAccess(
		models2.ManagerRole, models2.UnitEntity), controllers2.UpdateUnitAction)
	api.Delete("/unit/:id", middleware2.ResourceAccess(
		models2.ManagerRole, models2.UnitEntity), controllers2.DeleteUnitAction)

	api.Post("/maintenance", controllers2.CreateMaintenanceAction)
	api.Get("/maintenance/:id", controllers2.GetMaintenanceAction)
	api.Put("/maintenance/:id", controllers2.UpdateMaintenanceAction)
	api.Delete("/maintenance/:id", controllers2.DeleteMaintenanceAction)

	api.Post("/resolution", controllers2.CreateResolutionAction)
	api.Get("/resolution/:id", controllers2.GetResolutionAction)
	api.Put("/resolution/:id", controllers2.UpdateResolutionAction)
	api.Delete("/resolution/:id", controllers2.DeleteResolutionAction)

	api.Post("/vote", controllers2.CreateVoteAction)
	api.Get("/vote/:id", controllers2.GetVoteAction)
	api.Put("/vote/:id", controllers2.UpdateVoteAction)
	api.Delete("/vote/:id", controllers2.DeleteVoteAction)

	api.Post("/complaint", controllers2.CreateComplaintAction)
	api.Get("/complaint/:id", controllers2.GetComplaintAction)       // user or manager can view complaints.
	api.Put("/complaint/:id", controllers2.UpdateComplaintAction)    // user or manager can update complaints.
	api.Delete("/complaint/:id", controllers2.DeleteComplaintAction) // manager can delete

	api.Post("/file", controllers2.UploadFileAction)
	api.Get("/file/:id", controllers2.GetFileAction)
	api.Put("/file/:id", controllers2.UpdateFileAction)
	api.Delete("/file/:id", controllers2.DeleteFileAction)
}
