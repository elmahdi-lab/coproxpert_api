package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/repositories"
	"log/slog"
	"strings"
)

func IsOwnerMiddleware(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Next()
	}
	userUuid, err := uuid.Parse(id)

	resourceType := getResourceType(c.Path())

	slog.Info("Resource", " type: ", resourceType)

	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return unauthorizedResponse2(c)
	}

	userRepo := repositories.NewUserRepository()
	user, err := userRepo.FindByToken(authHeader)

	if err != nil || user.IsTokenExpired() || user.IsLocked() {
		return unauthorizedResponse2(c)
	}

	// check if the user is the owner of the resource:
	isOwner := user.ID == userUuid

	if !isOwner {
		return unauthorizedResponse2(c)
	}

	return c.Next()
}

// unauthorizedResponse is a helper function to send unauthorized response
func unauthorizedResponse2(c *fiber.Ctx) error {
	remoteIp := c.IP()
	requestedUrl := string(c.Request().RequestURI())

	slog.Error("Unauthorized request, ", "url: ", requestedUrl, "ip: ", remoteIp)

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Unauthorized",
	})
}

func getResourceType(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) >= 3 {
		return parts[2] // Assumes resource type is the third part of the path (e.g., /api/user/:id)
	}
	return ""
}
