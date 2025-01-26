package services

import (
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
)

func CountUnitsByUser(userID uuid.UUID) int64 {
	permissionRepository := repositories.NewPermissionRepository()
	if permissionRepository == nil {
		return 0
	}

	return permissionRepository.CountUnitsByUserIDAndEntity(userID)
}

func CountUnitGroupsByUser(userID uuid.UUID) int64 {
	permissionRepository := repositories.NewPermissionRepository()
	if permissionRepository == nil {
		return 0
	}

	return permissionRepository.CountUnitGroupsByUserIDAndEntity(userID)
}

func CreatePermission(userID uuid.UUID, entityID uuid.UUID, role models.Role, entityName models.EntityName) *models.Permission {
	permissionRepository := repositories.NewPermissionRepository()
	permission := &models.Permission{
		UserID:     userID,
		EntityID:   entityID,
		EntityName: entityName,
		Role:       role,
	}
	err := permissionRepository.Create(permission)
	if err != nil {
		return nil
	}
	return permission
}

func DeletePermission(userID uuid.UUID, entityID uuid.UUID) error {
	permissionRepository := repositories.NewPermissionRepository()
	return permissionRepository.DeleteByUserIDAndEntityID(userID, entityID)
}

func HasPermission(user *models.User, entityID uuid.UUID, role models.Role) bool {
	if user.IsSuperAdmin() {
		return true
	}
	for _, permission := range user.Permissions {
		if permission.EntityID == entityID && permission.Role >= role {
			return true
		}
	}
	return false
}
