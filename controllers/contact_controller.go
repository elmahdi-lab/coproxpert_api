// controllers/contact_controller.go

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/services"
)

func CreateContactAction(c *fiber.Ctx) error {
	contact := new(models.Contact)

	if err := c.BodyParser(contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	createdContact, err := services.CreateContact(contact, c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Contact created successfully", "contact": createdContact})
}

func GetContactAction(c *fiber.Ctx) error {
	id := c.Params("id")
	contactUUID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	contact, err := services.GetContact(contactUUID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"contact": contact})
}

func UpdateContactAction(c *fiber.Ctx) error {
	contact := new(models.Contact)

	if err := c.BodyParser(contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updatedContact, err := services.UpdateContact(contact)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Contact updated successfully", "contact": updatedContact})
}

func DeleteContactAction(c *fiber.Ctx) error {
	id := c.Params("id")
	contactUUID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	deleted := services.DeleteContact(contactUUID)

	if !deleted {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Contact not found"})
	}

	return c.JSON(fiber.Map{"message": "Contact deleted successfully"})
}
