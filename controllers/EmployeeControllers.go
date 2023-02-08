package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/puneet105/jwt-go-gin-gorm-poc/database"
	"github.com/puneet105/jwt-go-gin-gorm-poc/models"
	"net/http"
)

func GetEmployees(c *gin.Context) {
	var user []models.Employee
	database.DB.Find(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func GetEmployeesByID(c *gin.Context){
	var user models.Employee
	err := database.DB.Where("id = ?", c.Param("id")).First(&user).Error
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"Record Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteEmployeeByID(c *gin.Context){
	var user models.Employee
	err := database.DB.Where("id = ?", c.Param("id")).First(&user).Error
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"Record Not Found"})
		return
	}

	database.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
