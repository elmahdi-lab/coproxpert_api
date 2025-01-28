package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ithumans.com/coproxpert/cmd"
	"ithumans.com/coproxpert/models"
)

type UnitGroupRepository struct {
	db *gorm.DB
}

func NewUnitGroupRepository() *UnitGroupRepository {
	db := cmd.GetDB()
	if db == nil {
		return nil
	}
	return &UnitGroupRepository{db: db}
}

// TODO: Pagination

func (ur *UnitGroupRepository) FindAll(page int, pageSize int) ([]models.UnitGroup, error) {
	var unitGroups []models.UnitGroup
	if err := ur.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&unitGroups).Error; err != nil {
		return nil, err
	}
	return unitGroups, nil
}

func (ur *UnitGroupRepository) FindByID(id uuid.UUID) (*models.UnitGroup, error) {
	var unitGroup models.UnitGroup
	if err := ur.db.First(&unitGroup, id).Error; err != nil {
		return nil, err
	}
	return &unitGroup, nil
}

func (ur *UnitGroupRepository) Create(unitGroup *models.UnitGroup) error {
	return ur.db.Create(unitGroup).Error
}

func (ur *UnitGroupRepository) Update(unitGroup *models.UnitGroup) error {
	return ur.db.Save(unitGroup).Error
}

func (ur *UnitGroupRepository) Delete(unitGroup *models.UnitGroup) error {
	return ur.db.Delete(unitGroup).Error
}

func (ur *UnitGroupRepository) DeleteByID(id uuid.UUID) bool {
	err := ur.db.Delete(&models.UnitGroup{}, id)
	if err != nil {
		return true
	}
	return false
}

func (ur *UnitGroupRepository) CountByUserID(id uuid.UUID) int64 {
	var count int64
	ur.db.Model(&models.UnitGroup{}).Where("owner_id = ?", id).Count(&count)
	return count
}
