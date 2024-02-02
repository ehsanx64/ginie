package blog

import (
	"fmt"
	"ginie/lib"
	"log"

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
	var p1 Post

	res := lib.DB.Preload("Tags").Find(&p1, 1)
	if res.Error != nil {
		panic(res.Error)
	}

	for i, t := range p1.Tags {
		fmt.Printf("%02d: Title: %s, Name: %s\n", i, t.Title, t.Name)
	}

	var p1FirstTag *Tag
	p1FirstTag = p1.Tags[0]

	var p4 Post
	res = lib.DB.Find(&p4, 4)
	if res.Error != nil {
		panic(res.Error)
	}

	log.Println("post id:", p4.ID)

	lib.DB.Model(&p4).Association("Tags").Append(
		&Tag{
			Name:  "dummy-tag",
			Title: "Dummy Tag",
		},
		p1FirstTag,
		&Tag{
			Name:  "dummy-tag-2",
			Title: "Dummy Tag 2",
		},
	)

	lib.DB.Model(&p1).Association("Tags").Clear()

	p4TagsCount := lib.DB.Model(&p4).Association("Tags").Count()
	log.Printf("p4 has %d tags\n", p4TagsCount)

	var p3 Post
	res = lib.DB.Find(&p3, 3)
	if res.Error != nil {
		panic(res.Error)
	}
	lib.DB.Model(&p3).Association("Tags").Delete(p3.Tags)

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
