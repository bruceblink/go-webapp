package main

import (
	"concurrency/controllers"
	"concurrency/models"

	"github.com/gin-gonic/gin"
)

func main() {

	ginInstance := gin.Default()
	models.ConnectDatabase()
	ginInstance.GET("/employees", controllers.FindEmployees)
	err := ginInstance.Run()
	if err != nil {
		return
	}

}
