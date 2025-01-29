// controllers/maintenance_controller.go

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/data/models"
	"ithumans.com/coproxpert/data/services"
)

func CreateMaintenanceAction(c *fiber.Ctx) error {
	maintenance := new(models.Maintenance)

	if err := c.BodyParser(maintenance); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	createdMaintenance, err := services.CreateMaintenance(maintenance)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(createdMaintenance)
}

func GetMaintenanceAction(c *fiber.Ctx) error {
	id := c.Params("id")
	maintenanceUUID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	maintenance, err := services.GetMaintenanceByID(maintenanceUUID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(maintenance)
}

func UpdateMaintenanceAction(c *fiber.Ctx) error {
	maintenance := new(models.Maintenance)

	if err := c.BodyParser(maintenance); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updatedMaintenance, err := services.UpdateMaintenance(maintenance)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(updatedMaintenance)
}

func DeleteMaintenanceAction(c *fiber.Ctx) error {
	id := c.Params("id")
	maintenanceUUID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	deleted := services.DeleteMaintenanceByID(maintenanceUUID)

	if deleted != true {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "maintenance not deleted"})
	}

	return c.JSON(fiber.Map{"message": "Maintenance deleted successfully"})
}
