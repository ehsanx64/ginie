package blog

import (
	"fmt"
	"ginie/lib"

	"github.com/gin-gonic/gin"
)

func RegisterBlogRoutes(router *gin.Engine) {
	// Define API endpoints in router group
	blogRoute := router.Group("/blog")
	{
		blogRoute.GET("/", index)
		blogRoute.GET("/:id", view)

		blogRoute.GET("/test", test)
		blogRoute.GET("/api", api)
	}
}

func test(c *gin.Context) {
	lib.RenderHTML(c, "blog/test", lib.Option{
		"message": "Welcome to blog",
	})
}

func api(c *gin.Context) {
	var out string = ""

	//many2many()

	lib.RenderHTML(c, "blog/test", lib.Option{
		"message": out,
	})
}

func index(c *gin.Context) {
	lib.RenderHTML(c, "blog/index", lib.Option{
		"posts": GetPosts(),
	})
}

func view(c *gin.Context) {
	id := c.Param("id")
	lib.RenderHTML(c, "blog/view", lib.Option{
		"title": fmt.Sprintf("Post ID %s", id),
		"content": `
			This is a the content for the post
		`,
	})
}
