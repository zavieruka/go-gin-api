package controllers

import (
	"go-gin-api/database"
	"go-gin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowStudents(c *gin.Context) {
	c.JSON(200, gin.H{
		"students": models.Students,
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

func CreateStudent(c *gin.Context) {
	var aluno models.Student
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}
