package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// Define static routes
	router.StaticFile("/", "./resources/home.html")
	router.Static("/assets", "./assets")
	router.StaticFS("/static", http.Dir("./static"))

	// Define API endpoints
	router.GET("/api/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to gin-app",
		})
	})

	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run() // By default listen and serve on 0.0.0.0:8080
}
