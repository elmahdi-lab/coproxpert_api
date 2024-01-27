package cmd

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"ithumans.com/coproxpert/config"
	"os"
)

func Start() {
	app := fiber.New()
	config.RegisterRoutes(app)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	address := host + ":" + port
	fmt.Printf("Server is starting on %s\n", address)

	err := app.Listen(address)
	if err != nil {
		fmt.Printf("Failed to start the server: %s\n", err)
		return
	}
}
