// controllers/vote_controller.go

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/services"
)

func CreateVoteAction(c *fiber.Ctx) error {
	vote := new(models.Vote)

	if err := c.BodyParser(vote); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	createdVote, err := services.CreateVote(vote)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(createdVote)
}

func GetVoteAction(c *fiber.Ctx) error {
	id := c.Params("id")
	voteUUID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	vote, err := services.GetVoteByID(voteUUID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(vote)
}

func UpdateVoteAction(c *fiber.Ctx) error {
	vote := new(models.Vote)

	if err := c.BodyParser(vote); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updatedVote, err := services.UpdateVote(vote)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(updatedVote)
}

func DeleteVoteAction(c *fiber.Ctx) error {
	id := c.Params("id")
	voteUUID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	deleted := services.DeleteVoteByID(voteUUID)

	if deleted != true {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "vote not deleted"})
	}

	return c.JSON(fiber.Map{"message": "Vote deleted successfully"})
}
