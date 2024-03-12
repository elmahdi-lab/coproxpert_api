package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/services"
)

func CreateBuildingAction(c *fiber.Ctx) error {
	building := new(models.Building)

	if err := c.BodyParser(building); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	createdBuilding, err := services.CreateBuilding(building)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Building created successfully", "building": createdBuilding})

}

func GetBuildingAction(c *fiber.Ctx) error {

	id := c.Params("id")
	buildingUuid, err := uuid.Parse(id)
	building, err := services.GetBuilding(buildingUuid)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"building": building})
}

func UpdateBuildingAction(c *fiber.Ctx) error {

	building := new(models.Building)

	if err := c.BodyParser(building); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updatedBuilding, err := services.UpdateBuilding(building)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Building updated successfully", "building": updatedBuilding})
}

func DeleteBuildingAction(c *fiber.Ctx) error {

	id := c.Params("id")
	buildingUuid, _ := uuid.Parse(id)
	deleted := services.DeleteBuilding(buildingUuid)

	if deleted == true {
		return c.JSON(fiber.Map{"message": "Building deleted successfully"})
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Building not found"})
}
