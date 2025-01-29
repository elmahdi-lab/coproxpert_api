package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	models2 "ithumans.com/coproxpert/data/models"
	"ithumans.com/coproxpert/data/services"
)

type PermissionRequest struct {
	UserID     uuid.UUID          `json:"user_id"`
	EntityID   uuid.UUID          `json:"entity_id"`
	EntityName models2.EntityName `json:"entity_name,omitempty"`
	Role       models2.Role       `json:"role,omitempty"`
}

func CreatePermissionAction(c *fiber.Ctx) error {
	var req PermissionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	// Can Only assign User and Manager roles.

	if req.Role != models2.UserRole && req.Role != models2.ManagerRole {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid role",
		})
	}

	// Get the authenticated user from context
	user := c.Locals("user").(*models2.User)

	// Check if the authenticated user is the owner of the entity
	if !services.IsOwner(user, req.EntityName, req.EntityID) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Only the owner can manage permissions",
		})
	}

	// Create the permission
	permission := services.CreatePermission(req.UserID, req.EntityID, req.Role, req.EntityName)
	if permission == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create permission",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(permission)
}

func DeletePermissionAction(c *fiber.Ctx) error {
	var req PermissionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	// Get the authenticated user from context
	user := c.Locals("user").(*models2.User)

	// Check if the authenticated user is the owner of the entity
	if !services.IsOwner(user, req.EntityName, req.EntityID) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Only the owner can manage permissions",
		})
	}

	// Delete the permission
	if err := services.DeletePermission(req.UserID, req.EntityID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete permission",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Permission deleted successfully",
	})
}
