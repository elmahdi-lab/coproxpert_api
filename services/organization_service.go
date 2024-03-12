package services

import (
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
)

func CreateOrganization(o *models.Organization) (*models.Organization, error) {
	organizationRepository, _ := repositories.NewOrganizationRepository()
	err := organizationRepository.Create(o)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func GetOrganization(id uuid.UUID) (*models.Organization, error) {
	organizationRepository, _ := repositories.NewOrganizationRepository()

	organization, err := organizationRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return organization, nil
}

func UpdateOrganization(o *models.Organization) (*models.Organization, error) {
	organizationRepository, _ := repositories.NewOrganizationRepository()

	err := organizationRepository.Update(o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func DeleteOrganization(id uuid.UUID) bool {
	organizationRepository, _ := repositories.NewOrganizationRepository()

	deleted := organizationRepository.DeleteByID(id)
	if deleted == true {
		return true
	}
	return false
}
