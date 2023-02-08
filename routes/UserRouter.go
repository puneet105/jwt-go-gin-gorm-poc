package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/puneet105/jwt-go-gin-gorm-poc/controllers"
	"github.com/puneet105/jwt-go-gin-gorm-poc/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.RequireAuth)
	incomingRoutes.GET("/employees", controllers.GetEmployees)
	incomingRoutes.GET("/employees/:id", controllers.GetEmployeesByID)
	incomingRoutes.DELETE("/employees/:id", controllers.DeleteEmployeeByID)
}
