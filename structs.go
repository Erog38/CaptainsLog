package main

import "github.com/jinzhu/gorm"

type Config struct {
	//Fully Qualified Domain Name
	FQDN string `json: "fqdn"`
	//Port on which to run the server
	Port string `json: "port"`
	//Blog database
	Database string `json: "database"`
	//Username for access
	Username string `json: "username"`
	//Password for user
	Password string `json:"password"`
	//Where to log the output
	Logfile string `json:"logfile"`
	//Where the HTML static files live
	HTML string `json:"html"`
}

type Post struct {
	gorm.Model
	Name string `json:"name, omitempty"`
	Body string `json:"body, omitempty"`
}

type RawSearch struct {
	Search string `form:"search" json: "search"`
}

type Response struct {
	Success bool   `json:"success"`
	Err     string `json:"error"`
}

type SearchResponse struct {
	Response
	Posts []Post `json:"facilities"`
}

type UpdateResponse struct {
	Response
}

type InsertResponse struct {
	Response
}

type DeleteResponse struct {
	Response
}
