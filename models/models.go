package models

import "gorm.io/gorm"

type Employee struct{
	gorm.Model
	//ID				uint		`json:"id" gorm:"primary_key"`
	Name			string		`json:"name"`
	ContactNumber	string		`json:"contact_number"`
	Address			string		`json:"address"`
	Email			string		`json:"email"`
	Password		string		`json:"password"`
	Token			string		`json:"token"`
}
