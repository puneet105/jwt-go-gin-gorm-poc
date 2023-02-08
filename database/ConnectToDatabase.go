package database

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)
var DB *gorm.DB

func ConnectDatabase() *gorm.DB{
	db, err := gorm.Open(sqlite.Open("employee.db"), &gorm.Config{})
	if err != nil{
		fmt.Println("Failed to connect to database: ", err)
	}
	fmt.Println("Connected to DB........!!")
	DB = db
	return DB
}
