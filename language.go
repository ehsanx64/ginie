package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Language struct {
	Name string `uri:"name" binding:"required"`
}

const cookieName = "locale"

var ts = map[string]string{
	"Ginie": "جینی",
	"Home":  "خانه",
	"About": "درباره",
}

var locale string = "en"

func xlate(key string) string {
	log.Println("xlate(", locale, ")")
	if locale == "fa" {
		if val, ok := ts[key]; ok {
			return val
		}
	}

	return key
}

func setXlate(lang string) {
	locale = lang
}

func setLanguage(c *gin.Context, locale string) {
	c.Set("locale", locale)
	c.SetCookie(cookieName, locale, 0, "/", "localhost", false, true)
}

func getLanguage(c *gin.Context) (string, error) {
	cookie, err := c.Cookie(cookieName)

	if err != nil {
		return "en", err
	}

	return cookie, nil
}

/*
** Language middleware
 */
func LanguageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale, _ := getLanguage(c)
		setXlate(locale)
		c.Set("locale", locale)
	}
}
