package controllers

import (
	"github.com/gin-gonic/gin"
)

func ShowStudents(c *gin.Context) {
	students := []string{"Alice", "Bob", "Charlie"}
	c.JSON(200, gin.H{
		"students": students,
	})
}

func Hello(c *gin.Context) {
	name, present := c.Params.Get("name")

	if !present {
		c.JSON(400, gin.H{
			"error": "Name parameter is required",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Hello " + name + "!",
	})

}
