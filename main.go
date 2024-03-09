package main

import (
	_ "ariga.io/atlas-provider-gorm/gormschema"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"ithumans.com/coproxpert/cmd"
	"ithumans.com/coproxpert/config"
	"ithumans.com/coproxpert/fixtures"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Failed to load .env file: %s\n", err)
	}

	app := fiber.New()
	config.RegisterRoutes(app)

	_, err = cmd.GetDB()

	if err != nil {
		return
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	address := host + ":" + port
	fmt.Printf("Server is starting on %s\n", address)

	if os.Getenv("ENV") == config.Development {
		fixtures.CreateUser()
	}

	err = app.Listen(address)
	if err != nil {
		fmt.Printf("Failed to start the server: %s\n", err)
		return
	}
}
