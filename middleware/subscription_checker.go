package middleware

import (
	"github.com/gofiber/fiber/v2"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/services"
)

func CheckSubscriptionLimit(limitType models.SubscriptionLimitType) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*models.User)

		// TODO: add support for super admin
		//if user.IsSuperAdmin() {
		//	slog.Info("Super admin bypassed the subscription limit", "userID:", user.ID)
		//	return c.Next()
		//}

		subscription := services.FindSubscriptionByUser(user)

		if subscription == nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "User has no subscription"})
		}

		if limitType == models.UnitLimit {
			unitCount := services.CountUnitsByUser(user)
			if unitCount >= models.SubscriptionTiers[subscription.SubscriptionType].UnitsLimit {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Unit limit reached"})
			}
		}

		if limitType == models.UnitGroupLimit {
			unitGroupCount := services.CountUnitGroupsByUser(user)
			if unitGroupCount >= models.SubscriptionTiers[subscription.SubscriptionType].UnitGroupsLimit {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Unit group limit reached"})
			}
		}

		return c.Next()
	}
}
