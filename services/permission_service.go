package services

import (
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
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

	return true
}
