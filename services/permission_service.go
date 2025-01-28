package services

import (
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
)

func IsOwner(user *models.User, resourceType models.EntityName, resourceID uuid.UUID) bool {

	switch resourceType {
	case models.UnitGroupEntity:
		unitGroupRepository := repositories.NewUnitGroupRepository()
		unitGroup, _ := unitGroupRepository.FindByID(resourceID)
		if unitGroup == nil {
			return false
		}
		return unitGroup.OwnerID == user.ID
	case models.UnitEntity:
		unitRepository := repositories.NewUnitRepository()
		unit, _ := unitRepository.FindByID(resourceID)
		if unit == nil {
			return false
		}
		return unit.OwnerID == user.ID
	}

	return false
}

func HasPermission(user *models.User, entity models.EntityName, resourceID uuid.UUID, requiredRole models.Role) bool {
	for _, permission := range user.Permissions {
		if permission.EntityName == entity && permission.EntityID == resourceID {
			if permission.Role >= requiredRole {
				return true
			}
		}
	}

	return false
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
