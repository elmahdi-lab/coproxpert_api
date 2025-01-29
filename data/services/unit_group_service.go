package services

import (
	"github.com/google/uuid"
	"ithumans.com/coproxpert/data/models"
	"ithumans.com/coproxpert/data/repositories"
)

func CreateUnitGroup(unitGroup *models.UnitGroup) (*models.UnitGroup, error) {
	unitGroupRepository := repositories.NewUnitGroupRepository()
	err := unitGroupRepository.Create(unitGroup)
	if err != nil {
		return nil, err
	}
	return unitGroup, nil

}

func GetUnitGroup(id uuid.UUID) (*models.UnitGroup, error) {
	unitGroupRepository := repositories.NewUnitGroupRepository()

	unitGroup, err := unitGroupRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return unitGroup, nil
}

func UpdateUnitGroup(unitGroup *models.UnitGroup) (*models.UnitGroup, error) {
	unitGroupRepository := repositories.NewUnitGroupRepository()

	err := unitGroupRepository.Update(unitGroup)
	if err != nil {
		return nil, err
	}

	return unitGroup, nil
}

func DeleteUnitGroup(id uuid.UUID) bool {
	unitGroupRepository := repositories.NewUnitGroupRepository()

	deleted := unitGroupRepository.DeleteByID(id)
	if deleted == true {
		return true
	}
	return false
}

func CountUnitGroupsByUser(id uuid.UUID) int64 {
	unitGroupRepository := repositories.NewUnitGroupRepository()
	return unitGroupRepository.CountByUserID(id)
}
