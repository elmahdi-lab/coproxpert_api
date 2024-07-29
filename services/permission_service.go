package services

import (
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
)

func HasPermission(user *models.User, entityType models.EntityType, entityID uuid.UUID, role models.Role) bool {
	for _, permission := range user.Permissions {
		if permission.Role == models.SuperAdminRole {
			return true
		}

		if permission.EntityType == entityType && permission.EntityID == entityID && (permission.Role == role || permission.Role == models.AdminRole) {
			return true
		}
	}

	if entityType == models.UnitEntity {
		organizationRepository := repositories.NewOrganizationRepository()
		organization, err := organizationRepository.FindByUnitID(entityID)

		if err != nil {
			return false
		}

		return HasPermission(user, models.OrganizationEntity, organization.ID, models.ManagerRole)
	}

	return false
}
