package services

import (
	"github.com/google/uuid"
	"ithumans.com/coproxpert/data/models"
	"ithumans.com/coproxpert/data/repositories"
)

func CreateMaintenance(m *models.Maintenance) (*models.Maintenance, error) {
	maintenanceRepository := repositories.NewMaintenanceRepository()
	if maintenanceRepository == nil {
		return nil, nil
	}

	err := maintenanceRepository.Create(m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func GetMaintenanceByID(id uuid.UUID) (*models.Maintenance, error) {
	maintenanceRepository := repositories.NewMaintenanceRepository()
	if maintenanceRepository == nil {
		return nil, nil
	}

	maintenance, err := maintenanceRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return maintenance, nil
}

func UpdateMaintenance(m *models.Maintenance) (*models.Maintenance, error) {
	maintenanceRepository := repositories.NewMaintenanceRepository()
	if maintenanceRepository == nil {
		return nil, nil
	}

	err := maintenanceRepository.Update(m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func DeleteMaintenanceByID(id uuid.UUID) bool {
	maintenanceRepository := repositories.NewMaintenanceRepository()
	if maintenanceRepository == nil {
		return false
	}

	return maintenanceRepository.DeleteByID(id)
}
