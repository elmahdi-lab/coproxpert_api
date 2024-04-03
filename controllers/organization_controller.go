// controllers/organization_controller.go

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/services"
)

func CreateOrganizationAction(c *fiber.Ctx) error {
	organization := new(models.Organization)

	if err := c.BodyParser(organization); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	createdOrganization, err := services.CreateOrganization(organization)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Organization created successfully", "organization": createdOrganization})
}

func GetOrganizationAction(c *fiber.Ctx) error {
	id := c.Params("id")
	orgUUID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	organization, err := services.GetOrganization(orgUUID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"organization": organization})
}

func UpdateOrganizationAction(c *fiber.Ctx) error {
	organization := new(models.Organization)

	if err := c.BodyParser(organization); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updatedOrganization, err := services.UpdateOrganization(organization)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Organization updated successfully", "organization": updatedOrganization})
}

func DeleteOrganizationAction(c *fiber.Ctx) error {
	id := c.Params("id")
	orgUUID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	deleted := services.DeleteOrganization(orgUUID)

	if deleted {
		return c.JSON(fiber.Map{"message": "Organization deleted successfully"})
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Organization not found"})
}
