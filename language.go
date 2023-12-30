package main

import (
	"github.com/gin-gonic/gin"
)

const cookieName = "locale"

var ts = map[string]string{
	"Ginie": "جینی",
	"Home":  "خانه",
	"About": "درباره",
}

var locale string = "en"

func xlate(key string) string {
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
	c.SetCookie(cookieName, locale, 0, "/", "localhost", false, true)
}

func getLanguage(c *gin.Context) (string, error) {
	cookie, err := c.Cookie(cookieName)

	if err != nil {
		return "", err
	}

	return cookie, nil
}
