package main

import (
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

func postHandler(c *gin.Context) {
	id := c.Param("id")
	var post Post
	db.Where("id = ?", id).First(&post)

	c.HTML(http.StatusOK, "post.html", pongo2.Context{
		"Post": post,
	})
}
