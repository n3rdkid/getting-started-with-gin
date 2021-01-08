package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world!",
	})
}

func PostHomePage(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Hello world!",
	})
}

func QueryStrings(c *gin.Context) {
	name := c.Query("name")
	c.JSON(200, gin.H{
		"message": "Hello " + name + "!",
	})
}
func PathParameters(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"message": "Hello " + name + "!",
	})
}
func PostHandler(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Something went wrong!",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": string(value),
	})
}
func main() {
	r := gin.Default()
	r.GET("/", HomePage)
	r.POST("/", PostHomePage)
	r.GET("/query", QueryStrings)        // /query?name=saurav
	r.GET("/path/:name", PathParameters) // /path/saurav
	r.POST("/greet", PostHandler)
	r.Run()
}
