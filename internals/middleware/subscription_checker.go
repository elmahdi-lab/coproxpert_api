package middleware

import (
	"github.com/gofiber/fiber/v2"
	models2 "ithumans.com/coproxpert/data/models"
	services2 "ithumans.com/coproxpert/data/services"
)

func CheckSubscriptionLimit(limitType models2.SubscriptionLimitType) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*models2.User)

		subscription := services2.FindSubscriptionByUser(user)

		if subscription == nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "User has no subscription"})
		}

		if subscription.IsExpired() {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Subscription expired"})
		}

		if limitType == models2.UnitLimit {
			unitCount := services2.CountUnitsByUser(user.ID)
			if unitCount >= models2.SubscriptionTiers[subscription.SubscriptionType].UnitsLimit {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Unit limit reached"})
			}
		}

		if limitType == models2.UnitGroupLimit {
			unitGroupCount := services2.CountUnitGroupsByUser(user.ID)
			if unitGroupCount >= models2.SubscriptionTiers[subscription.SubscriptionType].UnitGroupsLimit {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Unit group limit reached"})
			}
		}

		return c.Next()
	}
}
