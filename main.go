package main

import (
	"github.com/gin-gonic/gin"
	"github.com/puneet105/jwt-go-gin-gorm-poc/database"
	"github.com/puneet105/jwt-go-gin-gorm-poc/routes"
)

func init(){
	db := database.ConnectDatabase()
	database.SyncDB(db)
}

func main(){
	route := gin.New()
	routes.AuthRoutes(route)
	routes.UserRoutes(route)
	route.Run()
}
