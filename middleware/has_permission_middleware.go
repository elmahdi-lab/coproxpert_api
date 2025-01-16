package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/services"
)

/*
	1. As a user I want to manage my profile, view my unit and unit group details, and view my organization details.
	2. As a manager I want to view and manage the allowed entities, example: add invoice, add maintenance, add inspection, view complaints.
	3. As an admin manage all the organization.
	4. As a super admin by pass all limitations.

	TODO: as a user, get access to the unit that belongs to the organization I have permissions for.
	TODO: as a manager or admin get access to the unit or resources that belongs to the organization with the correct Role Level.
*/

func HasPermission(entityType models.EntityType, role models.Role) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*models.User)
		uuID := uuid.Nil
		orgUuid, _ := uuid.Parse(c.Params("organizationID"))

		switch entityType {
		case models.UnitEntity:
			uuID, _ = uuid.Parse(c.Params("unitID"))

		case models.UnitGroupEntity:
			uuID, _ = uuid.Parse(c.Params("unitGroupID"))
		}

		if !services.HasPermission(user, entityType, uuID, role, orgUuid) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		return c.Next()

	}
}
