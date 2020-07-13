package main

import "github.com/gin-gonic/gin"
import "net/http"
import "time"


func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Welcome",
		})
	})
	router.GET("/daycounter", getDayCounter)
	router.POST("/daycounter", countTheDays)
	router.GET("/binary", getBinary)
	router.POST("/binary", postBinary)
	router.Run(":8080")
}

// day counters

func getDayCounter(c *gin.Context) {
	c.HTML(http.StatusOK, "daycounter.tmpl", gin.H{
		"title": "Day Counter",
	})
}

func countTheDays(c *gin.Context) {
	dayOne := c.PostForm("dayOne")
	t, err := time.Parse("2006-01-02 15:04:05", dayOne)
	dayTwo := c.PostForm("dayTwo")
	c.JSON(http.StatusOK, gin.H{"Day One":t, "Day Two" : dayTwo, "error": err})
}

// binary features
func getBinary(c *gin.Context) {
	c.HTML(http.StatusOK, "binaryconverter.tmpl", gin.H{
		"title": "Binary Converter",
		"displayNumber": "",
	})
}

func postBinary(c *gin.Context) {
	displayNumber := c.PostForm("number")
		c.HTML(http.StatusOK, "binaryconverter.tmpl", gin.H{
			"title": "Binary Converter",
			"displayNumber": displayNumber,
	})
}


