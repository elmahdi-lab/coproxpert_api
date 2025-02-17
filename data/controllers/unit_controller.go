// controllers/unit_controller.go

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/data/models"
	"ithumans.com/coproxpert/data/services"
	"ithumans.com/coproxpert/internals/events"
)

func CreateUnitAction(c *fiber.Ctx) error {

	user := c.Locals("user").(*models.User)
	unit := new(models.Unit)

	if err := c.BodyParser(unit); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	unit.OwnerID = user.ID

	createdUnit, err := services.CreateUnit(unit)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = events.PublishMessage(user.ID, createdUnit.ID, models.UnitEntity, events.Created)
	if err != nil {
		return err
	}

	return c.JSON(createdUnit)
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

	return c.JSON(unit)
}

func UpdateUnitAction(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	id := c.Params("id")
	parsedId, _ := uuid.Parse(id)
	if parsedId == uuid.Nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "id is required"})
	}

	unit := new(models.Unit)
	unit.ID = parsedId
	if err := c.BodyParser(unit); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error parsing body"})
	}

	updateUnit, err := services.UpdateUnit(unit)
	if err != nil {
		return err
	}

	err = events.PublishMessage(user.ID, updateUnit.ID, models.UnitEntity, events.Updated)
	if err != nil {
		return err
	}

	return c.JSON(updateUnit)
}

func DeleteUnitAction(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	id := c.Params("id")
	unitUUID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	deleted := services.DeleteUnitByID(unitUUID)

	if deleted != true {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Unit not deleted"})
	}

	err = events.PublishMessage(user.ID, unitUUID, models.UnitEntity, events.Deleted)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "Unit deleted successfully"})
}

func GetUnitsAction(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	units, err := services.GetUnitsByUser(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(units)
}
