package main

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	origins := os.Getenv("ORIGINS")

	app.Use(cors.New(cors.Config{
		AllowOrigins: origins,
	}))

	//app.Use(limiter.New(limiter.Config{
	//	Next: func(c *fiber.Ctx) bool {
	//		return c.IP() == "127.0.0.1"
	//	},
	//	Max:        20,
	//	Expiration: 30 * time.Second,
	//	KeyGenerator: func(c *fiber.Ctx) string {
	//		return c.Get("x-forwarded-for")
	//	},
	//	LimitReached: func(c *fiber.Ctx) error {
	//		return c.JSON(fiber.Map{"error": "Too many requests"})
	//	},
	//	// Storage: myCustomStorage{},
	//}))

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
