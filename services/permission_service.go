package services

import (
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
)

func HasPermission(userID uuid.UUID, entityType models.EntityType, entityID uuid.UUID, role models.Role) bool {
	permissionRepository := repositories.NewPermissionRepository()
	permissions, err := permissionRepository.FindByUserID(userID)
	if err != nil {
		return false
	}

	for _, permission := range permissions {
		if permission.Role == models.SuperAdminRole {
			return true
		}

		if permission.EntityType == entityType && permission.EntityID == entityID && permission.Role == role {
			return true
		}
	}

	// TODO: Try to minimize db queries by setting the entities once or find a better solution
	// If the entity type is a unit group, we need to check if the user is a manager of the organization
	if entityType == models.UnitGroupEntity {
		unitGroupRepository := repositories.NewUnitGroupRepository()
		unitGroup, err := unitGroupRepository.FindByID(entityID)
		if err != nil {
			return false
		}

		// Check if the user is an admin of the organization associated with the unit group
		return HasPermission(userID, models.OrganizationEntity, unitGroup.OrganizationID, models.AdminRole)
	}

	if entityType == models.UnitEntity {
		unitRepository := repositories.NewUnitRepository()
		unit, err := unitRepository.FindByID(entityID)
		if err != nil {
			return false
		}

		// Check if the user is a manager of the organization associated with the unit
		return HasPermission(userID, models.UnitGroupEntity, unit.UnitGroupID, models.ManagerRole)
	}

	return false
}
