package services

import (
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
)

func CreateBuilding(b *models.Building) (*models.Building, error) {
	buildingRepository, _ := repositories.NewBuildingRepository()
	err := buildingRepository.Create(b)
	if err != nil {
		return nil, err
	}
	return b, nil

}

func GetBuilding(id uuid.UUID) (*models.Building, error) {
	buildingRepository, _ := repositories.NewBuildingRepository()

	building, err := buildingRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return building, nil
}

func UpdateBuilding(b *models.Building) (*models.Building, error) {
	buildingRepository, _ := repositories.NewBuildingRepository()

	err := buildingRepository.Update(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func DeleteBuilding(id uuid.UUID) bool {
	buildingRepository, _ := repositories.NewBuildingRepository()

	deleted := buildingRepository.DeleteByID(id)
	if deleted == true {
		return true
	}
	return false
}
