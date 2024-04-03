package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
)

func CreateContact(contact *models.Contact, ctx *fiber.Ctx) (*models.Contact, error) {
	user := ctx.Locals("user").(*models.User)
	contact.UserID = user.ID

	contactRepository := repositories.NewContactRepository()
	err := contactRepository.Create(contact)
	if err != nil {
		return nil, err
	}
	return contact, nil
}

func UpdateContact(contact *models.Contact) (*models.Contact, error) {
	contactRepository := repositories.NewContactRepository()

	err := contactRepository.Update(contact)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

func GetContact(id uuid.UUID) (*models.Contact, error) {
	contactRepository := repositories.NewContactRepository()

	contact, err := contactRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return contact, nil
}

func GetAllContactsByUser(id uuid.UUID) ([]models.Contact, error) {
	contactRepository := repositories.NewContactRepository()
	contacts, err := contactRepository.FindByUserID(id)
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func DeleteContact(id uuid.UUID) bool {
	contactRepository := repositories.NewContactRepository()

	deleted := contactRepository.DeleteByID(id)
	return deleted
}
