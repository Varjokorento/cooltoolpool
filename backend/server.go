package main

import "github.com/gin-gonic/gin"
import "net/http"


func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	router.GET("/hello2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Kalle",
		})
	})
	router.Run(":8080")
}