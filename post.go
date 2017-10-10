package main

import (
	"html/template"
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func postHandler(c *gin.Context) {
	id := c.Param("id")
	var post Post
	db.Where("id = ?", id).First(&post)

	unsafe := blackfriday.MarkdownCommon([]byte(post.Body))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	c.HTML(http.StatusOK, "post.html", pongo2.Context{
		"Post":     post,
		"Markdown": template.HTML(html),
	})
}
