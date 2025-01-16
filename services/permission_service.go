package services

import (
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
)

func HasPermission(user *models.User, entityType models.EntityType, entityID uuid.UUID, role models.Role, orgUuid uuid.UUID) bool {
	// Check direct permissions
	for _, permission := range user.Permissions {
		if permission.Role == models.SuperAdminRole {
			return true
		}

		if permission.EntityType == entityType && permission.EntityID == entityID && permission.Role == role {
			return true
		}
	}

	// Initialize the repository
	organizationRepository := repositories.NewOrganizationRepository()
	organization, err := organizationRepository.FindByUnitID(orgUuid)

	var expectedRole models.Role

	switch entityType {
	case models.UnitEntity:
		if role == models.ManagerRole || role == models.AdminRole {
			expectedRole = role
		} else {
			expectedRole = models.ManagerRole
		}
	case models.UnitGroupEntity:
		expectedRole = models.AdminRole
	case models.OrganizationEntity:
		expectedRole = models.AdminRole
	}

	// Check if the organization was found
	if err != nil || organization == nil {
		return false
	}

	// Recursively check permissions at the organization level
	return HasPermission(user, models.OrganizationEntity, organization.ID, expectedRole, orgUuid)
}
