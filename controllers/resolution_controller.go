// controllers/resolution_controller.go

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/services"
)

func CreateResolutionAction(c *fiber.Ctx) error {
	resolution := new(models.Resolution)

	if err := c.BodyParser(resolution); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	createdResolution, err := services.CreateResolution(resolution)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Resolution created successfully", "resolution": createdResolution})
}

func GetResolutionAction(c *fiber.Ctx) error {
	id := c.Params("id")
	resolutionUUID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	resolution, err := services.GetResolutionByID(resolutionUUID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"resolution": resolution})
}

func UpdateResolutionAction(c *fiber.Ctx) error {
	resolution := new(models.Resolution)

	if err := c.BodyParser(resolution); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updatedResolution, err := services.UpdateResolution(resolution)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Resolution updated successfully", "resolution": updatedResolution})
}

func DeleteResolutionAction(c *fiber.Ctx) error {
	id := c.Params("id")
	resolutionUUID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	deleted := services.DeleteResolutionByID(resolutionUUID)

	if deleted != true {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Resolution deleted successfully"})
}
