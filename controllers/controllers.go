package controllers

import (
	"go-webapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateEmployeeInput struct {
	Name       string `json:"name" binding:"required"`
	Department string `json:"department" binding:"required"`
}

type UpdateEmployeeInput struct {
	Name       string `json:"name"`
	Department string `json:"department"`
}

func FindEmployees(context *gin.Context) {
	var employees []models.Employee
	models.DB.Find(&employees)
	context.JSON(http.StatusOK, gin.H{"data": employees})
}

func CreateEmployee(ctx *gin.Context) {
	var input CreateEmployeeInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	employee := models.Employee{Name: input.Name, Department: input.Department}
	models.DB.Create(&employee)
	ctx.JSON(http.StatusOK, gin.H{"data": employee})
}

func FindEmployee(ctx *gin.Context) {
	var employee models.Employee
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&employee).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": employee})
}

func UpdateEmployee(ctx *gin.Context) {
	var employee models.Employee
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&employee).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	var input UpdateEmployeeInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&employee).Updates(input)
	ctx.JSON(http.StatusOK, gin.H{"data": employee})
}

func DeleteEmployee(ctx *gin.Context) {
	var employee models.Employee
	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&employee).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	models.DB.Delete(&employee)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
