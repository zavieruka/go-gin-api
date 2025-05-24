package routes

import (
	"go-gin-api/controllers"
	"go-gin-api/models"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()

	models.Students = []models.Student{
		{Name: "Alice", Cpf: "12345678901", RG: "123456789"},
		{Name: "Bob", Cpf: "23456789012", RG: "234567890"},
		{Name: "Charlie", Cpf: "34567890123", RG: "345678901"},
	}

	r.GET("/students", controllers.ShowStudents)
	r.GET("/:name", controllers.Hello)

	r.Run(":8000")
}
