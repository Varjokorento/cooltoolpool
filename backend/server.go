package main

import (
	"backend/binaryConverter"
	"backend/dayCounter"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
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
	dayDifference := dayCounter.DayCount(dayOne, dayTwo)
	spew.Dump(dayDifference)
	var dateJson, _ = json.Marshal(dayDifference)
	spew.Dump(dateJson)
	c.JSON(http.StatusOK, gin.H{"Day Difference": dateJson})
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
