package services

import (
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
)

func CreateUnitGroup(b *models.UnitGroup) (*models.UnitGroup, error) {
	unitGroupRepository := repositories.NewUnitGroupRepository()
	err := unitGroupRepository.Create(b)
	if err != nil {
		return nil, err
	}
	return b, nil

}

func GetUnitGroup(id uuid.UUID) (*models.UnitGroup, error) {
	unitGroupRepository := repositories.NewUnitGroupRepository()

	unitGroup, err := unitGroupRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return unitGroup, nil
}

func UpdateUnitGroup(b *models.UnitGroup) (*models.UnitGroup, error) {
	unitGroupRepository := repositories.NewUnitGroupRepository()

	err := unitGroupRepository.Update(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func DeleteUnitGroup(id uuid.UUID) bool {
	unitGroupRepository := repositories.NewUnitGroupRepository()

	deleted := unitGroupRepository.DeleteByID(id)
	if deleted == true {
		return true
	}
	return false
}

func CountUnitGroupsByUser(user *models.User) int64 {
	unitGroupRepository := repositories.NewUnitGroupRepository()

	count := unitGroupRepository.CountUnitGroupsByUserID(user.ID)
	return count
}
