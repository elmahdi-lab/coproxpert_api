package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
)

func IsOwnerMiddleware(resourceType repositories.ResourceType) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*models.User)

		if user.isAdmin() {

		}

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
