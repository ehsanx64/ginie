package lib

import (
	"ginie/config"
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

func Xlate(key string) string {
	log.Println("xlate(", locale, ")")
	if locale == "fa" {
		if val, ok := ts[key]; ok {
			return val
		}
	}

	return key
}

func SetXlate(lang string) {
	locale = lang
}

func SetLanguage(c *gin.Context, locale string) {
	c.Set("locale", locale)
	c.SetCookie(cookieName, locale, 0, "/", config.DomainName, false, true)
}

func GetLanguage(c *gin.Context) (string, error) {
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
		locale, _ := GetLanguage(c)
		SetXlate(locale)
		c.Set("locale", locale)
	}
}
