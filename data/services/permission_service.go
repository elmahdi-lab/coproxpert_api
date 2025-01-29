package services

import (
	"github.com/google/uuid"
	models2 "ithumans.com/coproxpert/data/models"
	repositories2 "ithumans.com/coproxpert/data/repositories"
)

func IsOwner(user *models2.User, resourceType models2.EntityName, resourceID uuid.UUID) bool {

	switch resourceType {
	case models2.UnitGroupEntity:
		unitGroupRepository := repositories2.NewUnitGroupRepository()
		unitGroup, _ := unitGroupRepository.FindByID(resourceID)
		if unitGroup == nil {
			return false
		}
		return unitGroup.OwnerID == user.ID
	case models2.UnitEntity:
		unitRepository := repositories2.NewUnitRepository()
		unit, _ := unitRepository.FindByID(resourceID)
		if unit == nil {
			return false
		}
		return unit.OwnerID == user.ID
	}

	return false
}

func HasPermission(user *models2.User, entity models2.EntityName, resourceID uuid.UUID, requiredRole models2.Role) bool {
	for _, permission := range user.Permissions {
		if permission.EntityName == entity && permission.EntityID == resourceID {
			if permission.Role >= requiredRole {
				return true
			}
		}
	}

	return false
}

func CreatePermission(userID uuid.UUID, entityID uuid.UUID, role models2.Role, entityName models2.EntityName) *models2.Permission {
	permissionRepository := repositories2.NewPermissionRepository()
	permission := &models2.Permission{
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
	permissionRepository := repositories2.NewPermissionRepository()
	return permissionRepository.DeleteByUserIDAndEntityID(userID, entityID)
}
