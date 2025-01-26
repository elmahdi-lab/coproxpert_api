package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
)

func HasPermission(entityId uuid.UUID, role models.Role) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*models.User)
		if user.IsSuperAdmin() {
			return c.Next()
		}

		for _, userPermission := range user.Permissions {
			if userPermission.EntityID == entityId && userPermission.Role >= role {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Permission denied"})

	}
}
