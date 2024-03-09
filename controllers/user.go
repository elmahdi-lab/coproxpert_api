package controllers

import (
	"github.com/gofiber/fiber/v2"
	"ithumans.com/coproxpert/models"
)

func UserGreeting(c *fiber.Ctx) error {

	loggedUser := c.Locals("user").(*models.User)

	return c.JSON(fiber.Map{"message": "Hello User", "user": loggedUser})

}
