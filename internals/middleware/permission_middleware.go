package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/data/models"
	"ithumans.com/coproxpert/data/services"
)

func ResourceAccess(requiredRole models.Role, resourceType models.EntityName) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*models.User)
		resourceID, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid resource ID"})
		}

		// Check ownership first
		if services.IsOwner(user, resourceType, resourceID) {
			return c.Next()
		}

		// Check permissions
		if services.HasPermission(user, resourceType, resourceID, requiredRole) {
			return c.Next()
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
	}
}
