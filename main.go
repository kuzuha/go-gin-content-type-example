package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	UserFromForm struct {
		Name string `form:"name"`
	}
	UserFromJson struct {
		Name string `json:"name"`
	}
)

func router(g *gin.Engine) {
	g.Any("/any", func(c *gin.Context) {
		u := &UserFromJson{}
		err := c.Bind(u)
		if err != nil || u.Name == "" {
			u := &UserFromJson{}
			err := c.Bind(u)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"name": u.Name,
		})
	})
	g.POST("/json", func(c *gin.Context) {
		u := &UserFromJson{}
		err := c.Bind(u)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"name": u.Name,
		})
	})
	g.POST("/form", func(c *gin.Context) {
		u := &UserFromForm{}
		err := c.Bind(u)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"name": u.Name,
		})
	})
	g.GET("/query", func(c *gin.Context) {
		u := &UserFromForm{}
		err := c.Bind(u)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"name": u.Name,
		})
	})
}

func main() {
	g := gin.Default()
	router(g)
	g.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
