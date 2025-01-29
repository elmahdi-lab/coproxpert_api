package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	models2 "ithumans.com/coproxpert/data/models"
	"ithumans.com/coproxpert/data/services"
	"ithumans.com/coproxpert/internals/events"
)

func CreateUnitGroupAction(c *fiber.Ctx) error {
	user := c.Locals("user").(*models2.User)
	unitGroup := new(models2.UnitGroup)

	if err := c.BodyParser(unitGroup); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	unitGroup.OwnerID = user.ID
	createdUnitGroup, err := services.CreateUnitGroup(unitGroup)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = events.PublishMessage(user.ID, createdUnitGroup.ID, models2.UnitGroupEntity, events.Created)
	if err != nil {
		return err
	}

	return c.JSON(createdUnitGroup)

}

func GetUnitGroupAction(c *fiber.Ctx) error {

	id := c.Params("id")
	unitGroupUuid, err := uuid.Parse(id)
	unitGroup, err := services.GetUnitGroup(unitGroupUuid)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(unitGroup)
}

func UpdateUnitGroupAction(c *fiber.Ctx) error {
	user := c.Locals("user").(*models2.User)

	unitGroup := new(models2.UnitGroup)

	if err := c.BodyParser(unitGroup); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updatedUnitGroup, err := services.UpdateUnitGroup(unitGroup)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err = events.PublishMessage(updatedUnitGroup.ID, user.ID, models2.UnitGroupEntity, events.Updated)
	if err != nil {
		return err
	}
	return c.JSON(updatedUnitGroup)
}

func DeleteUnitGroupAction(c *fiber.Ctx) error {
	user := c.Locals("user").(*models2.User)

	id := c.Params("id")
	unitGroupUuid, _ := uuid.Parse(id)
	deleted := services.DeleteUnitGroup(unitGroupUuid)

	if deleted != true {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "UnitGroup not found"})
	}

	err := events.PublishMessage(unitGroupUuid, user.ID, models2.UnitGroupEntity, events.Deleted)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "UnitGroup deleted successfully"})

}
