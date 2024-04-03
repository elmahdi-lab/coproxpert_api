package controllers

import (
	"github.com/gofiber/fiber/v2"
	"ithumans.com/coproxpert/cmd"
)

func HealthCheck(ctx *fiber.Ctx) error {
	db := cmd.GetDB() // Get the database instance

	if db != nil {
		var result int
		db.Raw("SELECT 1").Scan(&result)
		if result == 1 {
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"healthcheck": "OK"})
		}
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"healthcheck": "FAILED"})
}
