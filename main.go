package main

import (
	"go-webapp/controllers"
	"go-webapp/models"

	"github.com/gin-gonic/gin"
)

func main() {

	ginInstance := gin.Default()
	models.ConnectDatabase()
	ginInstance.GET("/employees", controllers.FindEmployees)
	ginInstance.POST("/employees", controllers.CreateEmployee)
	ginInstance.GET("/employees/:id", controllers.FindEmployee)
	ginInstance.DELETE("/employees/:id", controllers.DeleteEmployee)
	ginInstance.PATCH("/employees/:id", controllers.UpdateEmployee)
	err := ginInstance.Run()
	if err != nil {
		return
	}

}
