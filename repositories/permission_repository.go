package repositories

//
//import (
//	"github.com/google/uuid"
//	"gorm.io/gorm"
//	"ithumans.com/coproxpert/cmd"
//	"ithumans.com/coproxpert/models"
//)
//
//type PermissionRepository struct {
//	db *gorm.DB
//}
//
//func NewPermissionRepository() (*PermissionRepository, error) {
//	db, err := cmd.GetDB()
//	if err != nil {
//		return nil, err
//	}
//	return &PermissionRepository{db: db}, nil
//}
//
//func (ur *PermissionRepository) FindByID(id uuid.UUID) (*models.Permission, error) {
//	var Permission models.Permission
//	if err := ur.db.First(&Permission, id).Error; err != nil {
//		return nil, err
//	}
//	return &Permission, nil
//}
//
//func (ur *PermissionRepository) Create(Permission *models.Permission) error {
//	return ur.db.Create(Permission).Error
//}
//
//func (ur *PermissionRepository) Update(Permission *models.Permission) error {
//	return ur.db.Save(Permission).Error
//}
//
//func (ur *PermissionRepository) Delete(Permission *models.Permission) error {
//	return ur.db.Delete(Permission).Error
//}
//
//func (ur *PermissionRepository) FindAll() ([]models.Permission, error) {
//	var Permissions []models.Permission
//	if err := ur.db.Find(&Permissions).Error; err != nil {
//		return nil, err
//	}
//	return Permissions, nil
//}
