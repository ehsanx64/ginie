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
	Tags    []*Tag `gorm:"many2many:post_tags;"`

	CategoryID uint
	Category   Category `gorm:"constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
}

type Tag struct {
	gorm.Model

	Name  string
	Title string
	Post  []*Post `gorm:"many2many:post_tags;"`
}

type Category struct {
	gorm.Model

	Name     string
	Title    string
	ParentId uint
}

var samplePosts = []Post{
	{
		Title:   "Post 1",
		Content: "Post 1 Contents",
		Tags: []*Tag{
			{
				Name:  "tag1",
				Title: "Tag 1",
			},
			{
				Name:  "tag2",
				Title: "Tag 2",
			},
		},
		Category: Category{
			Name:     "category1",
			Title:    "Category 1",
			ParentId: 0,
		},
	},
	{
		Title:   "Post 2",
		Content: "Post 2 Contents",
		Tags: []*Tag{
			{
				Name:  "tag3",
				Title: "Tag 3",
			},
			{
				Name:  "tag4",
				Title: "Tag 4",
			},
		},
	},
	{
		Title:   "Post 3",
		Content: "Post 3 Contents",
		Tags: []*Tag{
			{
				Name:  "tag5",
				Title: "Tag 5",
			},
			{
				Name:  "tag6",
				Title: "Tag 6",
			},
		},
	},
	{
		Title:   "Post 4",
		Content: "Post 4 Contents",
		Tags: []*Tag{
			{
				Name:  "tag7",
				Title: "Tag 7",
			},
			{
				Name:  "tag8",
				Title: "Tag 8",
			},
		},
	},
	{
		Title:   "Post 5",
		Content: "Post 5 Contents",
		Tags: []*Tag{
			{
				Name:  "tag9",
				Title: "Tag 9",
			},
			{
				Name:  "tag10",
				Title: "Tag 10",
			},
		},
	},
}

func SetupModel() {
	db := lib.DB

	var tables = []interface{}{
		&Post{},
		&Tag{},
		&Category{},
	}

	db.Migrator().DropTable("post_tags")
	for _, v := range tables {
		if db.Migrator().HasTable(v) {
			log.Println("Dropping the table...")
			db.Migrator().DropTable(v)
		}
	}

	// Migrate the schema
	db.AutoMigrate(&Post{}, &Tag{}, &Category{})

	/*
		firstCategory := &Category{
			Name:     "first",
			Title:    "First",
			ParentId: 0,
		}

		secondCategory := &Category{
			Name:     "second",
			Title:    "Second",
			ParentId: 0,
		}

		alphaCategory := &Category{
			Name:     "alpha",
			Title:    "Alpha",
			ParentId: 1,
		}
	*/

	// Check and if there are no records add some test items
	var posts []Post
	res := db.Find(&posts)
	if res.Error != nil {
		panic(res.Error)
	}

	if res.RowsAffected < 1 {
		log.Println("No posts found. Inserting some ...")
		db.Preload("Tag").Create(&samplePosts)
	} else {
		log.Println("No need to populate the posts table ...")
	}
}

func GetPosts() []Post {
	var posts []Post

	_ = lib.DB.Find(&posts)
	return posts
}
