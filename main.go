package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Option map[string]any
type NavLink map[string]string
type MainMenu map[int]NavLink

func setOptions(args Option) gin.H {
	out := gin.H{}

	for k, v := range args {
		out[k] = v
	}

	out["AppName"] = "Ginie"

	out["MainMenu"] = MainMenu{
		0: NavLink{
			"title": "Home",
			"url":   "/",
		},
		1: NavLink{
			"title": "RTL",
			"url":   "/rtl",
		},
		2: NavLink{
			"title": "About",
			"url":   "/about",
		},
	}

	fmt.Println(out)
	return out
}

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
		c.HTML(http.StatusOK, "pages/home.tmpl", setOptions(Option{
			"title": "Home",
		}))
	})

	router.GET("/rtl", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pages/home-fa.tmpl", setOptions(Option{
			"title": "Home",
			"rtl":   true,
		}))
	})

	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pages/about.tmpl", setOptions(Option{
			"title": "About",
		}))
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
