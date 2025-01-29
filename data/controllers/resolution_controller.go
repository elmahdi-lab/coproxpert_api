// controllers/resolution_controller.go

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/data/models"
	"ithumans.com/coproxpert/data/services"
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

	return c.JSON(createdResolution)
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

	return c.JSON(resolution)
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

	return c.JSON(updatedResolution)
}

func DeleteResolutionAction(c *fiber.Ctx) error {
	id := c.Params("id")
	resolutionUUID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	deleted := services.DeleteResolutionByID(resolutionUUID)

	if deleted != true {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "resolution not deleted"})
	}

	return c.JSON(fiber.Map{"message": "Resolution deleted successfully"})
}
