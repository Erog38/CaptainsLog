package main

import (
	"net/http"
	"strconv"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

func createHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "editor.html", pongo2.Context{
		"new": true,
	})
}

func editHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var blogPost Post
	db.Where("id = ?", id).First(&blogPost)
	c.HTML(http.StatusOK, "editor.html", pongo2.Context{
		"Post": blogPost,
	})
}

func postsHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Param("page"))

	pageOpts := PageOpts{}
	c.BindQuery(&pageOpts)

	if pageOpts.PageSize == 0 {
		pageOpts.PageSize = 25
	}

	offset := page * pageOpts.PageSize

	var blogPosts []Post
	db.Offset(offset).Limit(pageOpts.PageSize).Select("name, id").Find(&blogPosts)

	c.HTML(http.StatusOK, "list.html", pongo2.Context{
		"Posts": blogPosts,
	})
}

func updateHandler(c *gin.Context) {
	var post Post
	id := c.Param("id")
	url, _ := c.GetQuery("url")
	c.BindQuery(post)

	if post.Name == "" {
		c.HTML(http.StatusNotAcceptable, "confirmation.html", pongo2.Context{
			"body":    "your post needs a name!",
			"success": false,
			"url":     url,
		})
	}
	if post.Body == "" {
		c.HTML(http.StatusNotAcceptable, "confirmation.html", pongo2.Context{
			"body":    "your post needs a body!",
			"success": false,
			"url":     url,
		})
	}
	//else we have a name, body, and ID, continue!
	var blogPost Post
	db.Where("id = ?", id).First(&blogPost)

	blogPost.Name = post.Name
	blogPost.Body = post.Body
	db.Save(blogPost)

	c.HTML(http.StatusOK, "confirmation.html", pongo2.Context{
		"success": true,
		"url":     url,
		"id":      id,
	})
}

func insertHandler(c *gin.Context) {
}
