package main

import (
	"ginie/lib"
	"ginie/modules/api"
	"ginie/modules/blog"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(lib.LanguageMiddleware())

	router.SetFuncMap(template.FuncMap{
		"xlate": lib.Xlate,
	})

	// Define templates
	router.LoadHTMLGlob("templates/**/*")

	// Define static routes
	router.StaticFile("/landing", "./resources/landing.html")
	router.Static("/assets", "./assets")
	router.StaticFS("/static", http.Dir("./static"))

	// Define routes
	router.GET("/", func(c *gin.Context) {
		lib.RenderHTML(c, "pages/home", lib.Option{
			"title": lib.Xlate("Home"),
		})
	})

	router.GET("/about", func(c *gin.Context) {
		lib.RenderHTML(c, "pages/about", lib.Option{
			"title": lib.Xlate("About"),
		})
	})

	router.GET("/language/:name", func(c *gin.Context) {
		var l lib.Language
		if err := c.ShouldBindUri(&l); err != nil {
			c.JSON(400, gin.H{
				"msg": err,
			})

			return
		}

		lib.SetLanguage(c, l.Name)
		c.Redirect(http.StatusFound, "/")
	})

	api.RegisterApiRoutes(router)
	lib.InitDatabase()
	blog.RegisterBlogRoutes(router)
	blog.SetupModel()

	router.Run() // By default listen and serve on 0.0.0.0:8080
}
