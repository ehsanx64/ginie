package lib

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Option map[string]any
type NavLink map[string]string
type MainMenu map[int]NavLink

func SetOptions(args Option) gin.H {
	out := gin.H{}

	for k, v := range args {
		out[k] = v
	}

	out["AppName"] = "Ginie"
	out["MainMenu"] = MainMenu{
		0: NavLink{
			"title": Xlate("Home"),
			"url":   "/",
		},
		1: NavLink{
			"title": Xlate("About"),
			"url":   "/about",
		},
		2: NavLink{
			"title": Xlate("Blog"),
			"url":   "/blog",
		},
	}

	return out
}

func RenderHTML(c *gin.Context, template string, options Option) {
	locale := c.GetString("locale")

	rtl := false
	templ := template

	if locale == "fa" {
		rtl = true
		templ = templ + "-fa.tmpl"
	} else {
		templ = templ + ".tmpl"
	}

	opts := Option{
		"rtl": rtl,
	}

	for k, v := range options {
		opts[k] = v
	}

	log.Println("Active template:", templ)
	log.Println("Options:", opts)
	c.HTML(http.StatusOK, templ, SetOptions(opts))
}
