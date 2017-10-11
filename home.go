package main

import (
	"net/http"
	"strconv"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

func homeHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Param("page"))

	pageOpts := PageOpts{}
	c.BindQuery(&pageOpts)

	if pageOpts.PageSize == 0 {
		pageOpts.PageSize = 10
	}

	offset := page * pageOpts.PageSize

	var blogPosts []Post
	db.Offset(offset).Limit(pageOpts.PageSize).Find(&blogPosts)

	var summerizedPosts []Post
	for _, p := range blogPosts {
		if len(p.Body) > 100 {
			p.Body = p.Body[:100] + "..."
		}
		summerizedPosts = append(summerizedPosts, p)
	}
	var count int
	db.Model(&Post{}).Count(&count)

	var buttons []int

	for i := 0; i <= int(count/10); i++ {
		buttons = append(buttons, i)
	}

	c.HTML(http.StatusOK, "home.html", pongo2.Context{
		"Posts":   summerizedPosts,
		"Buttons": buttons,
		"Page":    page,
	})
}
