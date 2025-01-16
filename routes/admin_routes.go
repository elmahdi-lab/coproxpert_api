package routes

import (
	"github.com/gofiber/fiber/v2"
	"ithumans.com/coproxpert/controllers"
	"ithumans.com/coproxpert/middleware"
	"ithumans.com/coproxpert/models"
)

func RegisterAdminRoutes(app *fiber.App) {

	// Group with authentication middleware for secure routes
	api := app.Group("/api", middleware.AuthMiddleware)

	const userPath = "/user/:id"
	const organizationPath = "/:organizationID/organization"
	const unitGroupPath = "/:organizationID/unit-group/:id"
	const unitPath = "/:organizationID/:unitGroupID/unit/:id"
	const maintenancePath = "/:organizationID/maintenance/:id"
	const resolutionPath = "/:organizationID/resolution/:id"
	const votePath = "/:organizationID/:resolutionID/vote/:id"
	const complaintPath = "/:organizationID/complaint/:id"
	const filePath = "/:organizationID/file/:id"

	api.Post("/user", controllers.CreateUserAction)
	api.Get(userPath, controllers.GetUserAction)
	api.Put(userPath, controllers.UpdateUserAction)
	api.Delete(userPath, controllers.DeleteUserAction)
	api.Put("/user/password/:id", controllers.UpdatePasswordAction)
	api.Post("/user/logout", controllers.LogoutAction)

	api.Post("/user/subscribe/:type", controllers.Subscribe)

	//organizationEndpoints := api.Group("/organization")

	//organizationEndpoints.Post("", middleware.CheckSubscriptionLimit(models.OrganizationLimit), controllers.CreateOrganizationAction) // TODO: add a subscription limit check for organizations.
	//organizationEndpoints.Get(organizationPath, middleware.HasPermission(models.OrganizationEntity, models.ManagerRole), controllers.GetOrganizationAction)
	//organizationEndpoints.Put(organizationPath, middleware.HasPermission(models.OrganizationEntity, models.AdminRole), controllers.UpdateOrganizationAction)
	//organizationEndpoints.Delete(organizationPath, middleware.HasPermission(models.OrganizationEntity, models.ManagerRole), controllers.DeleteOrganizationAction)

	// Check subscription to limit the number of unit groups the user can create.
	api.Post("/unit-group", middleware.CheckSubscriptionLimit(models.UnitGroupLimit), controllers.CreateUnitGroupAction)
	api.Get(unitGroupPath, controllers.GetUnitGroupAction)
	api.Put(unitGroupPath, controllers.UpdateUnitGroupAction)
	api.Delete(unitGroupPath, controllers.DeleteUnitGroupAction)

	// Check subscription to limit the number of units the user can create.
	api.Post("/:unitGroupID/unit", middleware.CheckSubscriptionLimit(models.UnitLimit), controllers.CreateUnitAction)
	api.Get(unitPath, controllers.GetUnitAction)
	api.Put(unitPath, controllers.UpdateUnitAction)
	api.Delete(unitPath, controllers.DeleteUnitAction)

	api.Post("/maintenance", controllers.CreateMaintenanceAction)
	api.Get(maintenancePath, controllers.GetMaintenanceAction)
	api.Put(maintenancePath, controllers.UpdateMaintenanceAction)
	api.Delete(maintenancePath, controllers.DeleteMaintenanceAction)

	api.Post("/resolution", controllers.CreateResolutionAction)
	api.Get(resolutionPath, controllers.GetResolutionAction)
	api.Put(resolutionPath, controllers.UpdateResolutionAction)
	api.Delete(resolutionPath, controllers.DeleteResolutionAction)

	// TODO: allow only unit users to vote.
	api.Post("/:resolutionID/vote", controllers.CreateVoteAction)
	api.Get(votePath, controllers.GetVoteAction)
	api.Put(votePath, controllers.UpdateVoteAction)
	api.Delete(votePath, controllers.DeleteVoteAction)

	// TODO: allow only unit users to create complaints.
	api.Post("/complaint", controllers.CreateComplaintAction)
	api.Get(complaintPath, controllers.GetComplaintAction)       // user or manager can view complaints.
	api.Put(complaintPath, controllers.UpdateComplaintAction)    // user or manager can update complaints.
	api.Delete(complaintPath, controllers.DeleteComplaintAction) // manager can delete

	api.Post("/file", controllers.UploadFileAction)
	api.Get(filePath, controllers.GetFileAction)
	api.Put(filePath, controllers.UpdateFileAction)
	api.Delete(filePath, controllers.DeleteFileAction)
}
