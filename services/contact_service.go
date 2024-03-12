package services

import (
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
)

func CreateContact(c *models.Contact) (*models.Contact, error) {
	contactRepository, _ := repositories.NewContactRepository()
	err := contactRepository.Create(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func GetContact(id uuid.UUID) (*models.Contact, error) {
	contactRepository, _ := repositories.NewContactRepository()

	contact, err := contactRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return contact, nil
}

func UpdateContact(c *models.Contact) (*models.Contact, error) {
	contactRepository, _ := repositories.NewContactRepository()

	err := contactRepository.Update(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func DeleteContact(id uuid.UUID) bool {
	contactRepository, _ := repositories.NewContactRepository()

	deleted := contactRepository.DeleteByID(id)
	if deleted == true {
		return true
	}
	return false
}
