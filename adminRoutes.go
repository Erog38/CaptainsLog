package main

import (
	"net/http"
	"strconv"
	"strings"

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

	c.HTML(http.StatusOK, "posts.html", pongo2.Context{
		"Posts": blogPosts,
	})
}

func updateHandler(c *gin.Context) {
	id := c.Param("id")
	body, _ := c.GetPostForm("body")
	name, _ := c.GetPostForm("name")
	post := Post{Name: name, Body: body}
	if (strings.TrimSpace(post.Body) == "") && (strings.TrimSpace(post.Name) == "") {
		c.HTML(http.StatusBadRequest, "editor.html", pongo2.Context{
			"new":   true,
			"popup": "Your post needs a body and a name!",
			"Post":  post,
		})
		return
	}
	if strings.TrimSpace(post.Name) == "" {
		c.HTML(http.StatusBadRequest, "editor.html", pongo2.Context{
			"new":   true,
			"popup": "Your post needs a name!",
			"Post":  post,
		})
		return
	}
	if strings.TrimSpace(post.Body) == "" {
		c.HTML(http.StatusBadRequest, "editor.html", pongo2.Context{
			"new":   true,
			"popup": "Your post needs a body!",
			"Post":  post,
		})
		return
	}

	//else we have a name, body, and ID, continue!
	var blogPost Post
	db.Where("id = ?", id).First(&blogPost)

	blogPost.Name = post.Name
	blogPost.Body = post.Body
	db.Save(blogPost)
	c.HTML(http.StatusOK, "editor.html", pongo2.Context{
		"popup": "Updated!",
		"Post":  blogPost,
	})
}

func insertHandler(c *gin.Context) {
	body, _ := c.GetPostForm("body")
	name, _ := c.GetPostForm("name")
	post := Post{Name: name, Body: body}

	if (strings.TrimSpace(post.Body) == "") && (strings.TrimSpace(post.Name) == "") {
		c.HTML(http.StatusBadRequest, "editor.html", pongo2.Context{
			"new":   true,
			"popup": "Your post needs a body and a name!",
			"Post":  post,
		})
		return
	}
	if strings.TrimSpace(post.Name) == "" {
		c.HTML(http.StatusBadRequest, "editor.html", pongo2.Context{
			"new":   true,
			"popup": "Your post needs a name!",
			"Post":  post,
		})
		return
	}
	if strings.TrimSpace(post.Body) == "" {
		c.HTML(http.StatusBadRequest, "editor.html", pongo2.Context{
			"new":   true,
			"popup": "Your post needs a body!",
			"Post":  post,
		})
		return
	}
	db.Create(&post)
	c.HTML(http.StatusOK, "editor.html", pongo2.Context{
		"popup": "Posted!",
		"Post":  post,
	})
}

func deleteHandler(c *gin.Context) {
	id := c.Param("id")
	c.HTML(http.StatusOK, "delete.html", pongo2.Context{
		"ID": id,
	})
}
