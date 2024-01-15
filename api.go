package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(router *gin.Engine) {
	// Define API endpoints in router group
	apiRoute := router.Group("/api")
	{
		apiRoute.GET("/", index)
		apiRoute.GET("/ping", ping)

		apiRoute.GET("/form-data-bind/getb", formDataBind_getb)
		apiRoute.GET("/form-data-bind/getc", formDataBind_getc)
		apiRoute.GET("/form-data-bind/getd", formDataBind_getd)
	}
}

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to gin-app",
	})
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

type StructA struct {
	FieldA string `form:"field_a" binding:"required"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b" binding:"required"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}

type StructD struct {
	NestedAnonStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func formDataBind_getb(c *gin.Context) {
	var b StructB

	err := c.Bind(&b)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

func formDataBind_getc(c *gin.Context) {
	var sc StructC

	c.Bind(&sc)
	c.JSON(http.StatusOK, gin.H{
		"a": sc.NestedStructPointer,
		"c": sc.FieldC,
	})
}

func formDataBind_getd(c *gin.Context) {
	var d StructD

	c.Bind(&d)
	c.JSON(http.StatusOK, gin.H{
		"x": d.NestedAnonStruct,
		"d": d.FieldD,
	})
}
