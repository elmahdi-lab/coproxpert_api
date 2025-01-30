package main

import (
	"log/slog"
	"os"
	"time"

	_ "ariga.io/atlas-provider-gorm/gormschema"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"ithumans.com/coproxpert/cmd"
	"ithumans.com/coproxpert/internals/routes"
	"ithumans.com/coproxpert/pkg/config"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	err := godotenv.Load(".env")
	if err != nil {
	}

	timezone := os.Getenv("TIMEZONE")
	_, err = time.LoadLocation(timezone)
	if err != nil {
		return
	}
	app := fiber.New()

	app.Use(recover.New())
	routes.RegisterPublicRoutes(app)
	routes.RegisterUserRoutes(app)
	routes.RegisterAdminRoutes(app)

	_ = cmd.GetDB()

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	address := host + ":" + port
	logger.Info("Server is starting", "address", address)

	if os.Getenv("ENV") == config.Development {
		//fixtures.CreateUser()
	}

	err = app.Listen(address)
	if err != nil {
		logger.Error("Failed to start the server", "error", err)
		return
	}
}
