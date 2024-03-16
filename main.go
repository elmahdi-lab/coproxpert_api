package main

import (
	_ "ariga.io/atlas-provider-gorm/gormschema"
	"fmt"
	"github.com/gofiber/fiber/v2"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"ithumans.com/coproxpert/cmd"
	"ithumans.com/coproxpert/config"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Failed to load .env file: %s\n", err)
	}

	timezone := os.Getenv("TIMEZONE")
	_, err = time.LoadLocation(timezone)
	if err != nil {
		return
	}
	app := fiber.New()

	app.Use(recover2.New())
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
		//fixtures.CreateUser()
	}

	err = app.Listen(address)
	if err != nil {
		fmt.Printf("Failed to start the server: %s\n", err)
		return
	}
}
