package main

import (
	"net/http"
	"strconv"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

func createHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", pongo2.Context{})
}

func editHandler(c *gin.Context) {

}

func listHandler(c *gin.Context) {
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
