// controllers/complaint_controller.go

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/services"
)

func CreateComplaintAction(c *fiber.Ctx) error {
	complaint := new(models.Complaint)

	if err := c.BodyParser(complaint); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	createdComplaint, err := services.CreateComplaint(complaint)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Complaint created successfully", "complaint": createdComplaint})
}

func GetComplaintAction(c *fiber.Ctx) error {
	id := c.Params("id")
	complaintUUID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	complaint, err := services.GetComplaintByID(complaintUUID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"complaint": complaint})
}

func UpdateComplaintAction(c *fiber.Ctx) error {
	complaint := new(models.Complaint)

	if err := c.BodyParser(complaint); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updatedComplaint, err := services.UpdateComplaint(complaint)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Complaint updated successfully", "complaint": updatedComplaint})
}

func DeleteComplaintAction(c *fiber.Ctx) error {
	id := c.Params("id")
	complaintUUID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	deleted := services.DeleteComplaintByID(complaintUUID)

	if deleted != true {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Complaint deleted successfully"})
}
