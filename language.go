package main

import "github.com/gin-gonic/gin"

const cookieName = "locale"

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
