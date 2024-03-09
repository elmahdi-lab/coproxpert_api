package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"ithumans.com/coproxpert/repositories"
)

// AuthMiddleware is a middleware function for basic authentication
func AuthMiddleware(c *fiber.Ctx) error {
	// Get the "Authorization" header value
	authHeader := c.Get("Authorization")

	// Check if the header is empty
	if authHeader == "" {
		return unauthorizedResponse(c)
	}

	// find token in db
	tokenRepository, _ := repositories.NewTokenRepository()
	token, err := tokenRepository.FindByToken(authHeader, true)

	if err != nil || token.IsExpired() {
		return unauthorizedResponse(c)
	}

	c.Locals("user", token.User)

	return c.Next()
}

// unauthorizedResponse is a helper function to send unauthorized response
func unauthorizedResponse(c *fiber.Ctx) error {
	remoteIp := c.IP()
	requestedUrl := string(c.Request().RequestURI())

	fmt.Printf("Unauthorized request, url: %v, ip: %v\n", requestedUrl, remoteIp)
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Unauthorized",
	})
}
