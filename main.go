package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"ithumans.com/coproxpert/config"
	"ithumans.com/coproxpert/database"
	"ithumans.com/coproxpert/server"
	"ithumans.com/coproxpert/services/logging"
	"log"
	"os"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	env := os.Getenv("ENV")
	if env == config.Development {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	err = database.SetupDatabase()
	if err != nil {
		logger := logging.GetLogger()
		logger.LogError("Failed to connect to database")
		fmt.Printf("Failed to connect to database %s\n", err)
	}

	server.Start()

}
