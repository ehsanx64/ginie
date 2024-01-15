package main

import (
	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(router *gin.Engine) {
	// Define API endpoints in router group
	apiRoute := router.Group("/api")
	{
		apiRoute.GET("/", index)
		apiRoute.GET("/ping", ping)
	}
}

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to gin-app",
	})
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
