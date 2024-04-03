package services

import (
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
)

func CreateResolution(r *models.Resolution) (*models.Resolution, error) {
	resolutionRepository := repositories.NewResolutionRepository()
	if resolutionRepository == nil {
		return nil, nil
	}

	err := resolutionRepository.Create(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func GetResolutionByID(id uuid.UUID) (*models.Resolution, error) {
	resolutionRepository := repositories.NewResolutionRepository()
	if resolutionRepository == nil {
		return nil, nil
	}
	resolution, err := resolutionRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return resolution, nil
}

func UpdateResolution(r *models.Resolution) (*models.Resolution, error) {
	resolutionRepository := repositories.NewResolutionRepository()
	if resolutionRepository == nil {
		return nil, nil
	}

	err := resolutionRepository.Update(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func DeleteResolutionByID(id uuid.UUID) bool {
	resolutionRepository := repositories.NewResolutionRepository()
	if resolutionRepository == nil {
		return false
	}

	return resolutionRepository.DeleteByID(id)
}
