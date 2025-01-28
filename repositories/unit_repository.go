package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ithumans.com/coproxpert/cmd"
	"ithumans.com/coproxpert/models"
)

type UnitRepository struct {
	db *gorm.DB
}

func NewUnitRepository() *UnitRepository {
	db := cmd.GetDB()
	if db == nil {
		return nil
	}
	return &UnitRepository{db: db}
}

func (ur *UnitRepository) FindByUser(user *models.User) ([]models.Unit, error) {
	var unit []models.Unit
	if err := ur.db.Where("user_id = ?", user.ID).Find(&unit).Error; err != nil {
		return nil, err
	}
	return unit, nil
}

func (ur *UnitRepository) FindByID(id uuid.UUID) (*models.Unit, error) {
	var Unit models.Unit
	if err := ur.db.First(&Unit, id).Error; err != nil {
		return nil, err
	}
	return &Unit, nil
}

func (ur *UnitRepository) FindAll() ([]models.Unit, error) {
	var Units []models.Unit
	if err := ur.db.Find(&Units).Error; err != nil {
		return nil, err
	}
	return Units, nil
}

func (ur *UnitRepository) Create(Unit *models.Unit) error {
	return ur.db.Create(Unit).Error
}

func (ur *UnitRepository) Update(Unit *models.Unit) error {
	return ur.db.Save(Unit).Error
}

func (ur *UnitRepository) Delete(Unit *models.Unit) error {
	return ur.db.Delete(Unit).Error
}

func (ur *UnitRepository) DeleteByID(id uuid.UUID) bool {
	err := ur.db.Delete(&models.Unit{}, id)
	if err != nil {
		return true
	}
	return false
}

func (ur *UnitRepository) CountByUserID(id uuid.UUID) int64 {
	var count int64
	ur.db.Model(&models.Unit{}).Where("owner_id = ?", id).Count(&count)
	return count
}
