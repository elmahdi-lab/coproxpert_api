package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ithumans.com/coproxpert/cmd"
	"ithumans.com/coproxpert/models"
)

type PermissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository() *PermissionRepository {
	db := cmd.GetDB()
	if db != nil {
		return nil
	}
	return &PermissionRepository{db: db}
}

func (ur *PermissionRepository) FindByID(id uuid.UUID) (*models.Permission, error) {
	var Permission models.Permission
	if err := ur.db.First(&Permission, id).Error; err != nil {
		return nil, err
	}
	return &Permission, nil
}

func (ur *PermissionRepository) FindByUserID(userID uuid.UUID) ([]models.Permission, error) {
	var permissions []models.Permission
	if err := ur.db.Where("user_id = ?", userID).Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

func (ur *PermissionRepository) Create(Permission *models.Permission) error {
	return ur.db.Create(Permission).Error
}

func (ur *PermissionRepository) Update(Permission *models.Permission) error {
	return ur.db.Save(Permission).Error
}

func (ur *PermissionRepository) Delete(Permission *models.Permission) error {
	return ur.db.Delete(Permission).Error
}

func (ur *PermissionRepository) FindAll() ([]models.Permission, error) {
	var Permissions []models.Permission
	if err := ur.db.Find(&Permissions).Error; err != nil {
		return nil, err
	}
	return Permissions, nil
}

func (ur *PermissionRepository) FindByUserIDAndEntity(id uuid.UUID, entityType models.EntityType, entityID uuid.UUID) (*models.Permission, error) {
	var permission models.Permission
	if err := ur.db.Where("user_id = ? AND entity_type = ? AND entity_id = ?", id, entityType, entityID).Find(&permission).Error; err != nil {
		return nil, err
	}
	return &permission, nil
}
