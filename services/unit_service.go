package services

import (
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
)

func CreateUnit(u *models.Unit) (*models.Unit, error) {
	unitRepository := repositories.NewUnitRepository()
	if unitRepository == nil {
		return nil, nil
	}

	err := unitRepository.Create(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func GetUnitByID(id uuid.UUID) (*models.Unit, error) {
	unitRepository := repositories.NewUnitRepository()
	if unitRepository == nil {
		return nil, nil
	}
	unit, err := unitRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return unit, nil
}

func UpdateUnit(u *models.Unit) (*models.Unit, error) {

	unitRepository := repositories.NewUnitRepository()
	if u.ID == uuid.Nil {
		return nil, nil
	}

	oldUnit, err := GetUnitByID(u.ID)
	if err != nil {
		return nil, err
	}

	if oldUnit == nil {
		return nil, nil
	}

	if u.Name != "" {
		oldUnit.Name = u.Name
	}

	if u.Type != "" {
		oldUnit.Type = u.Type
	}

	if u.IsEnabled != nil {
		oldUnit.IsEnabled = u.IsEnabled
	}

	if unitRepository == nil {
		return nil, nil
	}
	err = unitRepository.Update(oldUnit)
	if err != nil {
		return nil, err
	}

	return oldUnit, nil
}

func DeleteUnitByID(id uuid.UUID) bool {
	unitRepository := repositories.NewUnitRepository()
	if unitRepository == nil {
		return false
	}

	return unitRepository.DeleteByID(id)
}
