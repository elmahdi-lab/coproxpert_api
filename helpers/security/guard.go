package security

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
)

func IsOwner(userID uuid.UUID, loggedUserID uuid.UUID) bool {
	if userID != loggedUserID {
		return false
	}
	return true
}

func Guard(ctx *fiber.Ctx, role models.Role) bool {
	user, ok := ctx.Locals("user").(*models.User)
	if !ok || user == nil {
		return false
	}

	if !userHasPermission(user, role) {
		return false
	}

	return true
}

func userHasPermission(user *models.User, role models.Role) bool {
	for _, permission := range user.Permissions {
		if permission.Role == role {
			return true
		}
	}
	return false
}
