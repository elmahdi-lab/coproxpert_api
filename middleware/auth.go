package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
	"ithumans.com/coproxpert/services"
)

type UserRepository interface {
	FindByToken(token string) (*models.User, error)
}

// AuthMiddleware is a middleware function for basic authentication
func AuthMiddleware(c *fiber.Ctx) error {
	// Get the "Authorization" header value
	authHeader := c.Get("Authorization")
	userRepo, _ := repositories.NewUserRepository()

	// Check if the header is empty
	if authHeader == "" {
		return unauthorizedResponse(c)
	}

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

	fmt.Printf("Unauthorized request, url: %v, ip: %v\n", requestedUrl, remoteIp)
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Unauthorized",
	})
}
