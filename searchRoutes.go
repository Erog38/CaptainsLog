package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func searchAPIHandler(c *gin.Context) {

	resp := SearchResponse{}
	pageOpts := PageOpts{}
	searchPost := Post{}

	c.BindJSON(&searchPost)
	c.BindQuery(&pageOpts)

	if pageOpts.PageSize == 0 {
		pageOpts.PageSize = 25
	}

	var blogPosts []Post

	offset := pageOpts.Page * pageOpts.PageSize
	if searchPost.isEmpty() {
		db.Offset(offset).Limit(pageOpts.PageSize).Find(&blogPosts)
	} else {
		db.Offset(offset).Limit(pageOpts.PageSize).Where(
			"name like ?", searchPost.Name).Or(
			"body like ?", searchPost.Body).Or(
			"id = ?", searchPost.ID).Find(&blogPosts)
	}
	resp.Success = true
	resp.Posts = blogPosts
	c.JSON(http.StatusOK, resp)
}
