package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func initRoutes(g *gin.Engine, config *Config) {

	g.GET("/", rootHandler)
	g.GET("/home", homeHandler)
	g.GET("/robots.txt", robotHandler)

	g.POST("/api/search", searchAPIHandler)
	g.POST("/api/search/raw", rawSearchAPIHandler)

	if strings.TrimSpace(config.Username) != "" &&
		strings.TrimSpace(config.Password) != "" {

		authorized := g.Group("/api", gin.BasicAuth(gin.Accounts{
			config.Username: config.Password,
		}))

		authorized.PUT("/update", updateAPIHandler)
		authorized.PUT("/insert", insertAPIHandler)
		authorized.DELETE("/delete", deleteAPIHandler)
	} else {

		g.PUT("/update", updateAPIHandler)
		g.PUT("/insert", insertAPIHandler)
		g.DELETE("/delete", deleteAPIHandler)
	}

}

//Handler functions
func rootHandler(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "/home")
}

func robotHandler(c *gin.Context) {
	c.String(http.StatusOK, "User-agent: *\nDisallow: /")
}
