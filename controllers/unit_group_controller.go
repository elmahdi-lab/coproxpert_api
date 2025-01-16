package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/helpers"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/services"
)

func CreateUnitGroupAction(c *fiber.Ctx) error {

	// TODO: Subscription is tied to the organization, so we need to set the org local
	//if helpers.IsSubscriptionLimitExceeded() == true {
	//	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Subscription limit exceeded"})
	//}

	user := c.Locals("user").(*models.User)
	unitGroup := new(models.UnitGroup)
	unitGroup.UserID = helpers.UuidPointer(user.ID)

	if err := c.BodyParser(unitGroup); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	createdUnitGroup, err := services.CreateUnitGroup(unitGroup)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "UnitGroup created successfully", "unitGroup": createdUnitGroup})

}

func GetUnitGroupAction(c *fiber.Ctx) error {

	id := c.Params("id")
	unitGroupUuid, err := uuid.Parse(id)
	unitGroup, err := services.GetUnitGroup(unitGroupUuid)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"unitGroup": unitGroup})
}

func UpdateUnitGroupAction(c *fiber.Ctx) error {

	unitGroup := new(models.UnitGroup)

	if err := c.BodyParser(unitGroup); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updatedUnitGroup, err := services.UpdateUnitGroup(unitGroup)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "UnitGroup updated successfully", "unitGroup": updatedUnitGroup})
}

func DeleteUnitGroupAction(c *fiber.Ctx) error {

	id := c.Params("id")
	unitGroupUuid, _ := uuid.Parse(id)
	deleted := services.DeleteUnitGroup(unitGroupUuid)

	if deleted == true {
		return c.JSON(fiber.Map{"message": "UnitGroup deleted successfully"})
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "UnitGroup not found"})
}
