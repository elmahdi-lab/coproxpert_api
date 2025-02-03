package middleware

import (
	"github.com/gofiber/fiber/v2"
	"ithumans.com/coproxpert/data/models"
	"ithumans.com/coproxpert/data/services"
)

func CheckSubscriptionLimit(limitType models.SubscriptionLimitType) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*models.User)

		subscription := services.FindSubscriptionByUser(user)

		if subscription == nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "User has no subscription"})
		}

		if subscription.IsExpired() {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Subscription expired"})
		}

		if limitType == models.UnitLimit {
			unitCount := services.CountUnitsByUser(user.ID)
			if unitCount >= models.SubscriptionTierConfigs[subscription.Tier].UnitsLimit {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Unit limit reached"})
			}
		}

		if limitType == models.UnitGroupLimit {
			unitGroupCount := services.CountUnitGroupsByUser(user.ID)
			if unitGroupCount >= models.SubscriptionTierConfigs[subscription.Tier].UnitGroupsLimit {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Unit group limit reached"})
			}
		}

		return c.Next()
	}
}
