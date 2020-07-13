package main

import (
	"backend/binaryConverter"
	"backend/dayCounter"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
	dayTwo := c.PostForm("dayTwo")
	dayOneAsInt, err := strconv.Atoi(dayOne)
	dayTwoAsInt, err := strconv.Atoi(dayTwo)
	if err == nil {
		dayThree := dayCounter.DayCount(dayOneAsInt, dayTwoAsInt)
		c.JSON(http.StatusOK, gin.H{"Day One": dayOne, "Day Two": dayThree})
	} else {
		c.JSON(http.StatusOK, gin.H{"Day One": "error", "Day Two": "error"})
	}
}

// binary features

func getBinary(c *gin.Context) {
	c.HTML(http.StatusOK, "binaryconverter.tmpl", gin.H{
		"title":         "Binary Converter",
		"displayNumber": "",
	})
}

func postBinary(c *gin.Context) {
	displayNumber := c.PostForm("number")
	displayNumberAsInt, _ := strconv.Atoi(displayNumber)
	displayNumberAsBinary := binaryConverter.ConvertToBinary(displayNumberAsInt)
	c.HTML(http.StatusOK, "binaryconverter.tmpl", gin.H{
		"title":         "Binary Converter",
		"displayNumber": displayNumberAsBinary,
	})
}
