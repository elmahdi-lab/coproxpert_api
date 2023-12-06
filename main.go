package main

import (
	"github.com/gin-gonic/gin"
	"ithumans.com/coproxpert/server"
	"os"
)

func main() {
	env := os.Getenv("ENVIRONMENT")

	if env == "" {
		env = gin.ReleaseMode
	}

	gin.SetMode(env)
	server.Start()

}
