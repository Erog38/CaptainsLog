package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/robvdl/pongo2gin"
	"github.com/thinkerou/favicon"
)

var db *gorm.DB
var config *Config
var defaultConfig = Config{
	Port:     "8080",
	Database: "./blog.db",
	Logfile:  "./blog.log",
	HTML:     "./templates",
}

func main() {

	//	gin.SetMode(gin.ReleaseMode)
	if len(os.Args) < 2 {
		config = &defaultConfig
	} else {
		config = parseConfig(os.Args[1])
	}
	var err error

	db, err = gorm.Open("sqlite3", config.Database+"?_busy_timeout=5000")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	db.AutoMigrate(&Post{})

	//f, err := os.Create(config.Logfile)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	g := gin.Default()

	g.Use(favicon.New("./static/favicon.ico"))

	g.HTMLRender = pongo2gin.Default()
	initRoutes(g, config)

	g.Run(config.FQDN + ":" + config.Port)
}

//Validation functions

func validateArgs(args []string) {

	if len(args) < 2 {
		fmt.Printf("%s <config filePath> \n", args[0])
		os.Exit(1)
	}
}

//Parsing functions

func parseConfig(filePath string) *Config {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	config := Config{}

	err = json.NewDecoder(f).Decode(&config)
	if err != nil {
		log.Fatalln(err)
	}

	return &config
}
