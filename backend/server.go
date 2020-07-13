package main

import "github.com/gin-gonic/gin"
import "net/http"


func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Welcome",
		})
	})
	router.GET("/daycounter", func(c *gin.Context) {
		c.HTML(http.StatusOK, "daycounter.tmpl", gin.H{
			"title": "Day Counter",
		})
	})
	router.GET("/binary", func(c *gin.Context) {
		c.HTML(http.StatusOK, "binaryconverter.tmpl", gin.H{
			"title": "Binary Converter",
			"displayNumber": "",
		})
	})
	router.POST("/binary", func(c *gin.Context) {
		displayNumber := c.PostForm("number")
		c.HTML(http.StatusOK, "binaryconverter.tmpl", gin.H{
			"title": "Binary Converter",
			"displayNumber": displayNumber,
		})
	})
	router.POST("/daycounter", func(c *gin.Context) {
		dayOne := c.PostForm("dayOne")
		dayTwo := c.PostForm("dayTwo")
		c.JSON(http.StatusOK, gin.H{"Day One":dayOne, "Day Two" : dayTwo})
	})
	router.Run(":8080")
}