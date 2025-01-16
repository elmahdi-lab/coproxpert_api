package controllers

import "github.com/gofiber/fiber/v2"

func LinkOrganizationToUser(c *fiber.Ctx) error {
	return c.SendString("Link organization to user")
}
