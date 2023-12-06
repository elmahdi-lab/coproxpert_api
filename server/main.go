package server

import (
	"fmt"
	"ithumans.com/coproxpert/config"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupServer() *gin.Engine {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	r := gin.Default()
	config.RegisterRoutes(r)

	return r
}

func Start() {
	r := setupServer()

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	err := r.Run(host + ":" + port)
	if err != nil {
		fmt.Printf("Failed to start the server: %s\n", err)
		return
	}
	fmt.Printf("Server is running on %s:%s\n", host, port)
}
