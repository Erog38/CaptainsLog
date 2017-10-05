package main

import "github.com/jinzhu/gorm"
import "strings"

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

func (p Post) isEmpty() bool {
	if strings.TrimSpace(p.Body) == "" {
		if strings.TrimSpace(p.Name) == "" {
			if p.ID == 0 {
				return true
			}
		}
	}
	return false
}

type PageOpts struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
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
	Posts []Post `json:"posts"`
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
