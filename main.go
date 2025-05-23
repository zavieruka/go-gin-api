package main

import (
	"github.com/gin-gonic/gin"
)

func showStudents(c *gin.Context) {
	students := []string{"Alice", "Bob", "Charlie"}
	c.JSON(200, gin.H{
		"students": students,
	})
}

func main() {
	r := gin.Default()

	r.GET("/students", showStudents)

	r.Run(":8000")
}
