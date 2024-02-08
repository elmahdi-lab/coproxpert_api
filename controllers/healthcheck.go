package controllers

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"ithumans.com/coproxpert/cmd"
)

func HealthCheck(c *fiber.Ctx) error {
	ctx := context.Background()

	client, err := cmd.GetClient(ctx)
	if err != nil {
		fmt.Printf("Failed to get database client: %s\n", err)
	}

	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err := tx.RawQuery(ctx, "SELECT 1").Exec()

	if err != nil {
		return err
	}

	if err == nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"healthcheck": "FAIL"})
	}
	return c.JSON(fiber.Map{"healthcheck": "OK"})
}
