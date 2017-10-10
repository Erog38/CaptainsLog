package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRoutes(g *gin.Engine, config *Config) {

	g.GET("/", rootHandler)
	g.GET("/home", rootHandler)
	g.Static("/static", "./static")
	g.GET("/home/:page", homeHandler)
	g.GET("/robots.txt", robotHandler)
	g.GET("/post/:id", postHandler)
	g.POST("/api/search", searchAPIHandler)

	admin := g.Group("/admin", gin.BasicAuth(
		gin.Accounts{config.Username: config.Password}))

	admin.GET("/create", createHandler)
	admin.POST("/insert", insertHandler)
	admin.POST("/update/:id", updateHandler)
	admin.GET("/edit/:id", editHandler)
	admin.GET("/posts", postsHandler)

	api := admin.Group("/api")
	api.POST("/update", updateAPIHandler)
	api.POST("/insert", insertAPIHandler)
	api.POST("/delete", deleteAPIHandler)
}

//Handler functions
func rootHandler(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/home/0")
}

func robotHandler(c *gin.Context) {
	c.String(http.StatusOK, "User-agent: *\nDisallow: /")
}
