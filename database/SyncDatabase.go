package database

import (
	"fmt"
	"github.com/puneet105/jwt-go-gin-gorm-poc/models"
	"gorm.io/gorm"
)

func SyncDB(db *gorm.DB) {
	err := db.AutoMigrate(&models.Employee{})
	if err != nil {
		fmt.Println("Failed to migrate employee table to database: ", err)
	}

	fmt.Println("Tables have been migrated to DB.....!!")
}