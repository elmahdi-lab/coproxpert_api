package controllers

import (
	"github.com/gofiber/fiber/v2"
	"ithumans.com/coproxpert/cmd"
)

func HealthCheck(ctx *fiber.Ctx) error {

	db, err := cmd.GetDB() // Get the database instance

	if err == nil {
		var result int
		db.Raw("SELECT 1").Scan(&result)

		if result == 1 {
			ctx.Status(fiber.StatusOK)
			return ctx.JSON(fiber.Map{"healthcheck": "OK"})
		}
	}
	ctx.Status(fiber.StatusInternalServerError)
	return ctx.JSON(fiber.Map{"healthcheck": "FAILED"})
}
