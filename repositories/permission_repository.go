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
	if db == nil {
		return nil
	}
	return &PermissionRepository{db: db}
}

func (pr *PermissionRepository) FindByID(id uuid.UUID) (*models.Permission, error) {
	var Permission models.Permission
	if err := pr.db.First(&Permission, id).Error; err != nil {
		return nil, err
	}
	return &Permission, nil
}

func (pr *PermissionRepository) FindByUserID(userID uuid.UUID) ([]models.Permission, error) {
	var permissions []models.Permission
	if err := pr.db.Where("user_id = ?", userID).Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

func (pr *PermissionRepository) Create(Permission *models.Permission) error {
	return pr.db.Create(Permission).Error
}

func (pr *PermissionRepository) Update(Permission *models.Permission) error {
	return pr.db.Save(Permission).Error
}

func (pr *PermissionRepository) Delete(Permission *models.Permission) error {
	return pr.db.Delete(Permission).Error
}

func (pr *PermissionRepository) FindAll() ([]models.Permission, error) {
	var Permissions []models.Permission
	if err := pr.db.Find(&Permissions).Error; err != nil {
		return nil, err
	}
	return Permissions, nil
}

func (pr *PermissionRepository) FindByUserIDAndEntity(id uuid.UUID, entityID uuid.UUID) (*models.Permission, error) {
	var permission models.Permission
	if err := pr.db.Where("user_id = ? AND entity_id = ?", id, entityID).Find(&permission).Error; err != nil {
		return nil, err
	}
	return &permission, nil
}

func (pr *PermissionRepository) DeleteByUserIDAndEntityID(userID uuid.UUID, entityID uuid.UUID) error {
	return pr.db.Where("user_id = ? AND entity_id = ?", userID, entityID).Delete(&models.Permission{}).Error
}
