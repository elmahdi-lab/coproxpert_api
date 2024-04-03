package middleware

import (
	"github.com/gofiber/fiber/v2"
	"ithumans.com/coproxpert/repositories"
	"ithumans.com/coproxpert/services"
	"log/slog"
)

// AuthMiddleware is a middleware function for basic authentication
func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return unauthorizedResponse(c)
	}

	userRepo := repositories.NewUserRepository()

	// find token in db
	user, err := userRepo.FindByToken(authHeader)

	if err != nil || user.IsTokenExpired() || user.IsLocked() {
		return unauthorizedResponse(c)
	}

	user.ExtendValidity()
	user, _ = services.UpdateUser(user)

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
