package services

import (
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
)

func CreateMaintenance(m *models.Maintenance) (*models.Maintenance, error) {
	maintenanceRepository, err := repositories.NewMaintenanceRepository()
	if err != nil {
		return nil, err
	}

	err = maintenanceRepository.Create(m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func GetMaintenanceByID(id uuid.UUID) (*models.Maintenance, error) {
	maintenanceRepository, err := repositories.NewMaintenanceRepository()
	if err != nil {
		return nil, err
	}

	maintenance, err := maintenanceRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return maintenance, nil
}

func UpdateMaintenance(m *models.Maintenance) (*models.Maintenance, error) {
	maintenanceRepository, err := repositories.NewMaintenanceRepository()
	if err != nil {
		return nil, err
	}

	err = maintenanceRepository.Update(m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func DeleteMaintenanceByID(id uuid.UUID) bool {
	maintenanceRepository, err := repositories.NewMaintenanceRepository()
	if err != nil {
		return false
	}

	return maintenanceRepository.DeleteByID(id)
}
