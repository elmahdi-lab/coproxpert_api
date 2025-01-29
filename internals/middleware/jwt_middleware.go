package middleware

import (
	"log/slog"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/data/repositories"
	"ithumans.com/coproxpert/internals/helpers/auth"
)

// JWTProtected is a middleware function for basic authentication
func JWTProtected(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return unauthorizedResponse(c)
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := auth.ValidateJWT(token)
	if err != nil {
		return unauthorizedResponse(c)
	}

	userID, _ := uuid.Parse(claims["sub"].(string))

	userRepo := repositories.NewUserRepository()
	user, err := userRepo.FindByID(userID)
	if err != nil || user == nil {
		return unauthorizedResponse(c)
	}

	if user.IsLocked() {
		return unauthorizedResponse(c)
	}

	c.Locals("user", user)
	return c.Next()
}

// unauthorizedResponse is a helper function to send unauthorized response
func unauthorizedResponse(c *fiber.Ctx) error {
	remoteIp := c.IP()
	requestedUrl := string(c.Request().RequestURI())

	slog.Error("Unauthorized request, ", "url: ", requestedUrl, "ip: ", remoteIp)

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Unauthorized",
	})
}
