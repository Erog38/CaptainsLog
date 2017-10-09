package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func initRoutes(g *gin.Engine, config *Config) {

	g.GET("/", rootHandler)
	g.GET("/home", rootHandler)
	g.Static("/static", "./static")
	g.GET("/home/:page", homeHandler)
	g.GET("/robots.txt", robotHandler)

	g.POST("/api/search", searchAPIHandler)

	if strings.TrimSpace(config.Username) != "" &&
		strings.TrimSpace(config.Password) != "" {

		admin := g.Group("/admin", gin.BasicAuth(
			gin.Accounts{
				config.Username: config.Password,
			}))
		admin.GET("/create", createHandler)
		admin.GET("/edit", editHandler)
		admin.GET("/list", listHandler)

		authorized := g.Group("/api", gin.BasicAuth(gin.Accounts{
			config.Username: config.Password,
		}))
		authorized.POST("/update", updateAPIHandler)
		authorized.POST("/insert", insertAPIHandler)
		authorized.POST("/delete", deleteAPIHandler)
	} else {

		g.POST("/update", updateAPIHandler)
		g.POST("/insert", insertAPIHandler)
		g.POST("/delete", deleteAPIHandler)
	}

}

//Handler functions
func rootHandler(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/home/0")
}

func robotHandler(c *gin.Context) {
	c.String(http.StatusOK, "User-agent: *\nDisallow: /")
}
