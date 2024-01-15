package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(LanguageMiddleware())

	router.SetFuncMap(template.FuncMap{
		"xlate": xlate,
	})

	// Define templates
	router.LoadHTMLGlob("templates/**/*")

	// Define static routes
	router.StaticFile("/landing", "./resources/landing.html")
	router.Static("/assets", "./assets")
	router.StaticFS("/static", http.Dir("./static"))

	// Define routes
	router.GET("/", func(c *gin.Context) {
		renderHTML(c, "pages/home", Option{
			"title": xlate("Home"),
		})
	})

	router.GET("/about", func(c *gin.Context) {
		renderHTML(c, "pages/about", Option{
			"title": xlate("About"),
		})
	})

	router.GET("/language/:name", func(c *gin.Context) {
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

	RegisterApiRoutes(router)

	router.Run() // By default listen and serve on 0.0.0.0:8080
}
