package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func updateAPIHandler(c *gin.Context) {

	blogPost := Post{}
	resp := UpdateResponse{}
	c.BindJSON(&blogPost)

	if blogPost.ID == 0 {
		resp.Success = false
		resp.Err = "1: ID must not be empty"
		c.JSON(http.StatusNotAcceptable, resp)
		return
	}
	if strings.Trim(blogPost.Body, "") == "" {
		resp.Success = false
		resp.Err = "2: body must not be empty"
		c.JSON(http.StatusNotAcceptable, resp)
		return
	}
	if strings.Trim(blogPost.Name, "") == "" {
		resp.Success = false
		resp.Err = "3: Name must not be empty"
		c.JSON(http.StatusNotAcceptable, resp)
		return
	}

	var oldPost Post
	db.First(&oldPost, blogPost.ID)
	oldPost.Name = blogPost.Name
	oldPost.Body = blogPost.Body
	db.Save(&oldPost)

	resp.Success = true
	c.JSON(http.StatusOK, resp)
}

func insertAPIHandler(c *gin.Context) {

	blogPost := Post{}
	resp := InsertResponse{}
	c.BindJSON(&blogPost)
	if strings.Trim(blogPost.Body, "") == "" {
		resp.Success = false
		resp.Err = "2: body must not be empty"
		c.JSON(http.StatusNotAcceptable, resp)
		return
	}
	if strings.Trim(blogPost.Name, "") == "" {
		resp.Success = false
		resp.Err = "3: Name must not be empty"
		c.JSON(http.StatusNotAcceptable, resp)
		return
	}

	var post Post

	db.Where("name = ?", blogPost.Name).First(&post)
	if post.Name != "" {
		resp.Success = false
		resp.Err = "post exists"
		c.JSON(http.StatusConflict, resp)
		return
	}

	db.Create(&blogPost)
	resp.Success = true
	c.JSON(http.StatusOK, resp)
}

func deleteAPIHandler(c *gin.Context) {
	blogPost := Post{}
	resp := DeleteResponse{}
	c.BindJSON(&blogPost)

	if blogPost.ID == 0 {
		resp.Success = false
		resp.Err = "1: ID must not be empty or zero"
		c.JSON(http.StatusNotAcceptable, resp)
		return
	}

	var oldPost Post
	db.First(&oldPost, blogPost.ID)
	if oldPost.ID == 0 {
		resp.Success = false
		resp.Err = "4: Post does not exist"
		c.JSON(http.StatusNotAcceptable, resp)
		return
	}

	db.Delete(&oldPost)
	resp.Success = true
	c.JSON(http.StatusOK, resp)
}
