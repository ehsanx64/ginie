package main

import (
	"fmt"
	"log"
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
		locale, _ := getLanguage(c)
		rtl := false
		templ := "pages/home.tmpl"

		if locale == "fa" {
			rtl = true
			templ = "pages/home-fa.tmpl"
		}

		log.Println("Active template:", templ)
		c.HTML(http.StatusOK, templ, setOptions(Option{
			"title": "Home",
			"rtl":   rtl,
		}))
	})

	router.GET("/about", func(c *gin.Context) {
		locale, _ := getLanguage(c)
		rtl := false
		templ := "pages/about.tmpl"

		if locale == "fa" {
			rtl = true
			templ = "pages/about-fa.tmpl"
		}

		log.Println("Active template:", templ)
		c.HTML(http.StatusOK, templ, setOptions(Option{
			"title": "About",
			"rtl":   rtl,
		}))
	})

	router.GET("/language/:name", func(c *gin.Context) {
		type Language struct {
			Name string `uri:"name" binding:"required"`
		}

		var l Language
		if err := c.ShouldBindUri(&l); err != nil {
			c.JSON(400, gin.H{
				"msg": err,
			})

			return
		}

		setLanguage(c, l.Name)
		c.Redirect(http.StatusFound, "/")
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
