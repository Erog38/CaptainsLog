# Captain's Log

This is a blog with a REST API built for CS497.

It currently requires the following packages to function:

* github.com/gin-gonic/gin
* github.com/mattn/go-sqlite3
* github.com/fatih/structs
* github.com/jinzhu/gorm
* github.com/robvdl/pongo2gin
* github.com/thinkerou/favicon
* github.com/flosch/pongo2
* github.com/russross/blackfriday
* github.com/microcosm-cc/bluemonday

## Building

For this project to build correctly, make sure it is properly located in your
$GOPATH and that you have the required dependencies listed above. Then use the
following command:

`go build `

## Running

The server takes an optional config file which can be used to provide the 
following data:

<pre>

{
    fqdn: "localhost",
    port: "8080",
    database: "./blog.db",
    username: "generalSkroob",
    password: "1234",
    logfile: "./blog.log",
    html: "./html"
}

</pre>

This is passed to the program as follows:

`./captainslog config.json` 

If no config file is given, the server will default to the following:

<pre>
{
    fqdn: "localhost",
    port: "8080",
    database: "./blog.db",
    html:    "./html"
}
</pre>
