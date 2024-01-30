package blog

import (
	"ginie/lib"
	"log"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model

	Title   string
	Content string
	Image   string
}

var samplePosts = []Post{
	{
		Title:   "Post 1",
		Content: "Post 1 Contents",
	},
	{
		Title:   "Post 2",
		Content: "Post 2 Contents",
	},
	{
		Title:   "Post 3",
		Content: "Post 3 Contents",
	},
	{
		Title:   "Post 4",
		Content: "Post 4 Contents",
	},
	{
		Title:   "Post 5",
		Content: "Post 5 Contents",
	},
}

func SetupModel() {
	db := lib.DB

	// Migrate the schema
	db.AutoMigrate(&Post{})

	// Check and if there are no records add some test items
	var posts []Post
	res := db.Find(&posts)
	if res.Error != nil {
		panic(res.Error)
	}

	if res.RowsAffected < 1 {
		log.Println("No posts found. Inserting some ...")
		db.Create(&samplePosts)
	} else {
		log.Println("No need to populate the blog table ...")
	}
}

func GetPosts() []Post {
	var posts []Post

	_ = lib.DB.Find(&posts)
	return posts
}
