package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ithumans.com/coproxpert/cmd"
	"ithumans.com/coproxpert/models"
)

type BuildingRepository struct {
	db *gorm.DB
}

func NewBuildingRepository() (*BuildingRepository, error) {
	db, err := cmd.GetDB()
	if err != nil {
		return nil, err
	}
	return &BuildingRepository{db: db}, nil
}

// TODO: Pagination

func (ur *BuildingRepository) FindAll(page int, pageSize int) ([]models.Building, error) {
	var buildings []models.Building
	if err := ur.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&buildings).Error; err != nil {
		return nil, err
	}
	return buildings, nil
}

func (ur *BuildingRepository) FindByID(id uuid.UUID) (*models.Building, error) {
	var building models.Building
	if err := ur.db.First(&building, id).Error; err != nil {
		return nil, err
	}
	return &building, nil
}

func (ur *BuildingRepository) Create(building *models.Building) error {
	return ur.db.Create(building).Error
}

func (ur *BuildingRepository) Update(building *models.Building) error {
	return ur.db.Save(building).Error
}

func (ur *BuildingRepository) Delete(building *models.Building) error {
	return ur.db.Delete(building).Error
}

func (ur *BuildingRepository) DeleteByID(id uuid.UUID) bool {
	err := ur.db.Delete(&models.Building{}, id)
	if err != nil {
		return true
	}
	return false
}
