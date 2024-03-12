package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"ithumans.com/coproxpert/cmd"
	"ithumans.com/coproxpert/models"
)

type PropertyRepository struct {
	db *gorm.DB
}

func NewPropertyRepository() (*PropertyRepository, error) {
	db, err := cmd.GetDB()
	if err != nil {
		return nil, err
	}
	return &PropertyRepository{db: db}, nil
}

func (ur *PropertyRepository) FindByID(id uuid.UUID) (*models.Property, error) {
	var property models.Property
	if err := ur.db.First(&property, id).Error; err != nil {
		return nil, err
	}
	return &property, nil
}

func (ur *PropertyRepository) Create(property *models.Property) error {
	return ur.db.Create(property).Error
}

func (ur *PropertyRepository) Update(property *models.Property) error {
	return ur.db.Save(property).Error
}

func (ur *PropertyRepository) Delete(property *models.Property) error {
	return ur.db.Delete(property).Error
}

func (ur *PropertyRepository) FindAll() ([]models.Property, error) {
	var properties []models.Property
	if err := ur.db.Find(&properties).Error; err != nil {
		return nil, err
	}
	return properties, nil
}

func (ur *PropertyRepository) DeleteByID(id uuid.UUID) bool {
	err := ur.db.Delete(&models.Property{}, id)
	if err != nil {
		return true
	}
	return false
}
