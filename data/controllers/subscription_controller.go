package controllers

import (
	"github.com/gofiber/fiber/v2"
	models2 "ithumans.com/coproxpert/data/models"
	"ithumans.com/coproxpert/data/services"
)

func Subscribe(c *fiber.Ctx) error {
	user := c.Locals("user").(*models2.User)
	subscriptionType := c.Params("type")

	subscription, err := services.CreateSubscription(user, models2.SubscriptionType(subscriptionType))
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
	//err := services.(user, models.SubscriptionType(subscriptionType))
	//if err != nil {
	//	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//		"message": err.Error(),
	//	})
	//}

	return c.JSON(fiber.Map{
		"message": "Subscription deleted successfully",
	})
}
