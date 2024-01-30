package blog

import (
	"fmt"
	"ginie/lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterBlogRoutes(router *gin.Engine) {
	// Define API endpoints in router group
	blogRoute := router.Group("/blog")
	{
		blogRoute.GET("/", index)
		blogRoute.GET("/:id", view)
	}

}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "blog/test", lib.SetOptions(lib.Option{
		"message": "Welcome to blog",
	}))
}

func view(c *gin.Context) {
	id := c.Param("id")
	c.HTML(http.StatusOK, "blog/view", lib.SetOptions(lib.Option{
		"title": fmt.Sprintf("Post ID %s", id),
		"content": `
			This is a the content for the post
		`,
	}))
}
