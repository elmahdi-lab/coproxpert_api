package server

import (
	"fmt"
	"ithumans.com/coproxpert/config"
	"ithumans.com/coproxpert/services/logging"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupServer() *gin.Engine {
	err := godotenv.Load()
	if err != nil {
		logger := logging.GetLogger()
		logger.LogError("Error loading .env file")
	}

	r := gin.Default()
	config.RegisterRoutes(r)

	return r
}

func Start() {
	r := setupServer()

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	// Print the server start message before calling r.Run()
	fmt.Printf("Server is starting on %s:%s\n", host, port)

	err := r.Run(host + ":" + port)
	if err != nil {
		logger := logging.GetLogger()
		logger.LogError("Error loading .env file")
		fmt.Printf("Failed to start the server: %s\n", err)
		return
	}
}
