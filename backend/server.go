package main

import (
	"backend/binaryConverter"
	"backend/dayCounter"
	"backend/structs"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

func main() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Cool Tool Pool",
		})
	})
	router.GET("/about", getAbout)
	router.GET("/daycounter", getDayCounter)
	router.POST("/daycounter", countTheDays)
	router.GET("/binary", getBinary)
	router.POST("/binary", postBinary)
	router.Run(":8080")
}

// markdown

func getAbout(c *gin.Context) {
	mdfile, _ := ioutil.ReadFile("./markdown/" + "about.md")
	const postName = "Placeholder"
	postHTML := template.HTML(blackfriday.MarkdownCommon([]byte(mdfile)))
	post := structs.Post{Title: postName, Content: postHTML}
	c.HTML(http.StatusOK, "post.tmpl.html", gin.H{
		"Title":   post.Title,
		"Content": post.Content,
	})
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
