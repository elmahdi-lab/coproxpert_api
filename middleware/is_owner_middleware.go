package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
)

// TODO: this middleware does not take into account if the user is an admin.

func IsOwnerMiddleware(resourceType repositories.ResourceType) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*models.User)
		id := c.Params("id")

		if id == "" {
			return c.Next()
		}
		resourceUuid, err := uuid.Parse(id)
		repository, exists := repositories.RepositoryMap[resourceType]

		if !exists {
			return unauthorizedResponse(c)
		}
		resource, err := repository.FindByIDAndUserID(resourceUuid, user.ID)

		if resource == nil || err != nil {
			return unauthorizedResponse(c)
		}
		c.Locals("resource", resource)
		return c.Next()
	}
}
