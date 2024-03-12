package security

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
)

func IsOwnerOrAdmin(c *fiber.Ctx, e *uuid.UUID) bool {

	u := c.Locals("user").(*models.User)

	if u == nil || e == nil {
		return false
	}

	// check if

	// Check if the user is an admin
	for _, p := range u.Permissions {
		if *p.Role == models.AdminRole && u.ID == *e {
			return true
		}

	}
	return false
}

func Guard(c *fiber.Ctx, u *models.User, r models.Role, t *models.EntityType, e *uuid.UUID) error {
	isGranted := checkPermission(u, r, t, e)

	if !isGranted {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	return nil
}

func checkPermission(u *models.User, r models.Role, t *models.EntityType, e *uuid.UUID) bool {
	// Check if the user or any of the parameters are nil
	if u == nil {
		return false
	}

	// Check if the user has permissions, if not, return false
	if u.Permissions == nil {
		return false
	}

	// if the user has r permission with the same access level, entity type and entity id, return true:
	for _, p := range u.Permissions {
		if *p.Role == r && *p.EntityType == *t && *p.EntityID == *e {
			return true
		}
	}

	// if the entity type is r building, and the user has r permission as r manager with the organization linked to the building, return true:
	if *t == models.BuildingEntity {
		buildingRepository, _ := repositories.NewBuildingRepository()
		building, _ := buildingRepository.FindByID(*e)

		for _, p := range u.Permissions {
			if *p.Role == models.ManagerRole && *p.EntityType == models.OrganizationEntity {
				if *p.EntityID == building.OrganizationID {
					return true
				}
			}
		}
	}

	// if the entity type is r property, and the user has r permission as r manager of the building linked to the property, return true:
	if *t == models.PropertyEntity {
		propertyRepository, _ := repositories.NewPropertyRepository()
		property, _ := propertyRepository.FindByID(*e)
		for _, p := range u.Permissions {
			if *p.Role == models.ManagerRole && *p.EntityType == models.BuildingEntity {
				if *p.EntityID == property.BuildingID {
					return true
				}
			}
		}
	}

	return false
}
