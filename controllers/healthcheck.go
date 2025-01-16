package controllers

import (
	"github.com/gofiber/fiber/v2"
	"ithumans.com/coproxpert/cmd"
)

func HealthCheck(ctx *fiber.Ctx) error {
	// Create a health status map
	healthStatus := fiber.Map{
		"database":  "FAILED",
		"g-storage": "FAILED",
		"g-pubsub":  "FAILED",
	}

	// Database health check
	db := cmd.GetDB() // Get the database instance
	if db != nil {
		var result int
		err := db.Raw("SELECT 1").Scan(&result).Error
		if err == nil && result == 1 {
			healthStatus["database"] = "OK"
		}
	}

	// Google Cloud Storage health check
	gstorage := cmd.TestStorageConnection(ctx.Context())
	if gstorage == true {
		healthStatus["g-storage"] = "OK"
	}

	gpubsub := cmd.TestPubSubConnection(ctx.Context())
	if gpubsub == true {
		healthStatus["g-pubsub"] = "OK"
	}

	for _, status := range healthStatus {
		if status != "OK" {
			return ctx.Status(fiber.StatusInternalServerError).JSON(healthStatus)
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(healthStatus)
}
