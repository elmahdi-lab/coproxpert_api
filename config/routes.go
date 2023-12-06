package config

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Define your routes here
	r.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })
}
