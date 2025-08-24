package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/gzip"
)

func main() {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.LoadHTMLGlob("templates/*")
	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	r.Static("/css", "./static/css")
	r.Static("/js", "./static/js")
	r.GET("/", home)
	r.Run()
}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":   "Home Page",
		"message": "Welcome to my Gin app!",
	})
}
