// controllers/unit_controller.go

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/services"
)

func CreateUnitAction(c *fiber.Ctx) error {
	unit := new(models.Unit)

	if err := c.BodyParser(unit); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	createdUnit, err := services.CreateUnit(unit)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Unit created successfully", "unit": createdUnit})
}

func GetUnitAction(c *fiber.Ctx) error {
	id := c.Params("id")
	unitUUID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	unit, err := services.GetUnitByID(unitUUID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"unit": unit})
}

func UpdateUnitAction(c *fiber.Ctx) error {
	unit := new(models.Unit)

	if err := c.BodyParser(unit); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updatedUnit, err := services.UpdateUnit(unit)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Unit updated successfully", "unit": updatedUnit})
}

func DeleteUnitAction(c *fiber.Ctx) error {
	id := c.Params("id")
	unitUUID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	deleted := services.DeleteUnitByID(unitUUID)

	if deleted != true {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Unit deleted successfully"})
}
