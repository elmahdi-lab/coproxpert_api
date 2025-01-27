package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// ExtractAndValidateID deprecated
func ExtractAndValidateID(c *fiber.Ctx) error {
	path := c.Route().Path
	if strings.Contains(path, ":id") != true {
		return c.Next()
	}

	unitID := c.Params("id")
	if unitID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "id is required"})
	}

	// Parse UUID
	parsedUnitID, err := uuid.Parse(unitID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id format"})
	}

	// Store the parsed UUID in the context's local variables
	c.Locals("uuid", parsedUnitID)

	// Proceed to the next handler
	return c.Next()
}
