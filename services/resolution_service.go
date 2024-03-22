package services

import (
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
)

func CreateResolution(r *models.Resolution) (*models.Resolution, error) {
	resolutionRepository, err := repositories.NewResolutionRepository()
	if err != nil {
		return nil, err
	}

	err = resolutionRepository.Create(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func GetResolutionByID(id uuid.UUID) (*models.Resolution, error) {
	resolutionRepository, err := repositories.NewResolutionRepository()
	if err != nil {
		return nil, err
	}

	resolution, err := resolutionRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return resolution, nil
}

func UpdateResolution(r *models.Resolution) (*models.Resolution, error) {
	resolutionRepository, err := repositories.NewResolutionRepository()
	if err != nil {
		return nil, err
	}

	err = resolutionRepository.Update(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func DeleteResolutionByID(id uuid.UUID) bool {
	resolutionRepository, err := repositories.NewResolutionRepository()
	if err != nil {
		return false
	}

	return resolutionRepository.DeleteByID(id)
}
