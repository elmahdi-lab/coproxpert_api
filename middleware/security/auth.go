// middleware/security/security.go

package security

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware is a middleware function for basic authentication
func AuthMiddleware(c *fiber.Ctx) error {
	// Get the "Authorization" header value
	authHeader := c.Get("Authorization")

	// Check if the header is empty
	if authHeader == "" {
		// If empty, return unauthorized
		remoteIp := c.IP()
		requestedUrl := string(c.Request().RequestURI())

		fmt.Printf("Unauthorized request, url: %v, ip: %v\n", requestedUrl, remoteIp)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Here, you may implement your own logic to validate the credentials.
	// For simplicity, let's assume a basic authentication format "Basic base64(username:password)"
	// You should decode the base64 string and validate the username and password.

	// For example:
	// authParts := strings.Split(authHeader, " ")
	// if len(authParts) != 2 || authParts[0] != "Basic" {
	//     return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	//         "message": "Invalid authorization format",
	//     })
	// }
	// decoded, err := base64.StdEncoding.DecodeString(authParts[1])
	// if err != nil {
	//     return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	//         "message": "Invalid base64 encoding",
	//     })
	// }
	// credentials := strings.Split(string(decoded), ":")
	// if len(credentials) != 2 || !isValidUser(credentials[0], credentials[1]) {
	//     return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	//         "message": "Invalid username or password",
	//     })
	// }

	// Uncomment and customize the above code according to your authentication logic.

	// If the credentials are valid, proceed to the next handler
	return c.Next()
}
