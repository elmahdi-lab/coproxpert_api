// controllers/maintenance_controller.go

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/services"
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

	return c.JSON(fiber.Map{"message": "Maintenance created successfully", "maintenance": createdMaintenance})
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

	return c.JSON(fiber.Map{"maintenance": maintenance})
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

	return c.JSON(fiber.Map{"message": "Maintenance updated successfully", "maintenance": updatedMaintenance})
}

func DeleteMaintenanceAction(c *fiber.Ctx) error {
	id := c.Params("id")
	maintenanceUUID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	deleted := services.DeleteMaintenanceByID(maintenanceUUID)

	if deleted != true {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Maintenance deleted successfully"})
}
