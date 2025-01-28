package middleware

import (
	"github.com/gofiber/fiber/v2"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/services"
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
			if unitCount >= models.SubscriptionTiers[subscription.SubscriptionType].UnitsLimit {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Unit limit reached"})
			}
		}

		if limitType == models.UnitGroupLimit {
			unitGroupCount := services.CountUnitGroupsByUser(user.ID)
			if unitGroupCount >= models.SubscriptionTiers[subscription.SubscriptionType].UnitGroupsLimit {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Unit group limit reached"})
			}
		}

		return c.Next()
	}
}
