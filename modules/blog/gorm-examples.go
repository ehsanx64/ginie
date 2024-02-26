package blog

import (
	"fmt"
	"ginie/lib"
	"log"
)

func many2many() {
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
}

func simple() {

}
