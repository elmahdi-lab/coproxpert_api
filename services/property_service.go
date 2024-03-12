package services

import (
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
)

func CreateProperty(p *models.Property) (*models.Property, error) {
	propertyRepository, _ := repositories.NewPropertyRepository()
	err := propertyRepository.Create(p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func GetProperty(id uuid.UUID) (*models.Property, error) {
	propertyRepository, _ := repositories.NewPropertyRepository()

	property, err := propertyRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return property, nil
}

func UpdateProperty(p *models.Property) (*models.Property, error) {
	propertyRepository, _ := repositories.NewPropertyRepository()

	err := propertyRepository.Update(p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func DeleteProperty(id uuid.UUID) bool {
	propertyRepository, _ := repositories.NewPropertyRepository()

	deleted := propertyRepository.DeleteByID(id)
	if deleted == true {
		return true
	}
	return false
}
