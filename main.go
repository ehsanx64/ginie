package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Define templates
	router.LoadHTMLGlob("templates/**/*")

	// Define static routes
	router.StaticFile("/landing", "./resources/landing.html")
	router.Static("/assets", "./assets")
	router.StaticFS("/static", http.Dir("./static"))

	// Define routes
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pages/home.tmpl", gin.H{
			"title": "Home",
		})
	})

	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pages/about.tmpl", gin.H{
			"title": "About",
		})
	})

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
