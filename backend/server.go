package main

import (
	"backend/binaryConverter"
	"backend/dayCounter"
	"backend/gpaCounter"
	"backend/structs"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
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
	router.GET("/gpacounter", getGpaCounter)
	router.POST("/gpacounter", postGPACounter)
	router.POST("/daycounter", countTheDays)
	router.GET("/binary", getBinary)
	router.POST("/binary", postBinary)
	router.Run(":8080")
}

// gpaCounter
func getGpaCounter(c *gin.Context) {
	c.HTML(http.StatusOK, "gpaCounter.tmpl", gin.H{
		"title": "GPA Counter",
	})
}

func postGPACounter(c *gin.Context) {
	totalCredits := c.PostForm("credits")
	sumOfGrade := c.PostForm("grades")
	totalCreditsInt, _ := strconv.Atoi(totalCredits)
	sumOfGradeInt, _ := strconv.Atoi(sumOfGrade)
	gpa := gpaCounter.GetGpa(totalCreditsInt, sumOfGradeInt)
	spew.Dump(gpa)
	gpaJSON, _ := json.Marshal(gpa)
	spew.Dump(gpaJSON)
	c.Data(http.StatusOK, "application/json", gpaJSON)

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
	c.Data(http.StatusOK, "application/json", dateJson)
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
