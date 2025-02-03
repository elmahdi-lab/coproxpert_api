package controllers

import (
	"github.com/gofiber/fiber/v2"
	"ithumans.com/coproxpert/data/models"
	"ithumans.com/coproxpert/data/services"
)

func Subscribe(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	subscriptionType := c.Params("type")

	subscription, err := services.CreateSubscription(user, models.SubscriptionTier(subscriptionType))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(subscription)
}

func Unsubscribe(c *fiber.Ctx) error {
	//user := c.Locals("user").(*models.User)
	//subscriptionType := c.Params("type")
	//
	//err := services.(user, models.Tier(subscriptionType))
	//if err != nil {
	//	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//		"message": err.Error(),
	//	})
	//}

	return c.JSON(fiber.Map{
		"message": "Subscription deleted successfully",
	})
}
