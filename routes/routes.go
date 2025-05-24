package routes

import (
	"go-gin-api/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()

	r.GET("/students", controllers.ShowStudents)
	r.GET("/:name", controllers.Hello)

	r.Run(":8000")
}
