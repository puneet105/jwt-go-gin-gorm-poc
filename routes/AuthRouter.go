package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/puneet105/jwt-go-gin-gorm-poc/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/signup", controllers.SignUp)
	incomingRoutes.POST("/login", controllers.Login)
}
