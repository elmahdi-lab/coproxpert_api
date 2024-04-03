package services

import (
	"errors"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
)

func CreateComplaint(c *models.Complaint) (*models.Complaint, error) {
	complaintRepository := repositories.NewComplaintRepository()
	if complaintRepository == nil {
		return nil, errors.New("error creating complaint repository")
	}

	err := complaintRepository.Create(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func GetComplaintByID(id uuid.UUID) (*models.Complaint, error) {
	complaintRepository := repositories.NewComplaintRepository()
	if complaintRepository == nil {
		return nil, errors.New("error creating complaint repository")
	}

	complaint, err := complaintRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return complaint, nil
}

func UpdateComplaint(c *models.Complaint) (*models.Complaint, error) {
	complaintRepository := repositories.NewComplaintRepository()
	if complaintRepository == nil {
		return nil, errors.New("error creating complaint repository")
	}

	err := complaintRepository.Update(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func DeleteComplaintByID(id uuid.UUID) bool {
	complaintRepository := repositories.NewComplaintRepository()
	if complaintRepository == nil {
		return false
	}

	return complaintRepository.DeleteByID(id)
}
