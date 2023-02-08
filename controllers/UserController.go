package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/puneet105/jwt-go-gin-gorm-poc/database"
	"github.com/puneet105/jwt-go-gin-gorm-poc/helpers"
	"github.com/puneet105/jwt-go-gin-gorm-poc/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)
func SignUp(c *gin.Context){
	//get email and password off the body
	var user models.Employee

	if c.Bind(&user) != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//generate hash of password
	hashPassword, err := helpers.GeneratePasswordHash(c, user.Password)
	if err != nil{
		fmt.Println("Error generating hash password: ", err)
	}
	user.Password = hashPassword
	//create user in database
	//user := models.Employee{Name: body.Name, ContactNumber: body.ContactNumber, Address: body.Address, Email: body.Email, Password: hashPassword}
	tokenString := helpers.GenerateJwtToken(c, user.Name, user.ContactNumber, user.Address, user.Email)
	user.Token = tokenString
	result := database.DB.Create(&user)
	if result.Error != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Failed to create user in database",
		})
		return
	}
	//send response
	c.JSON(http.StatusOK, gin.H{
		"data" : "User has successfully signed up",
	})

}

func Login(c *gin.Context){
	//get email and password off the body
	var user models.Employee

	if c.Bind(&user) != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	//look for requested user
	var foundUser models.Employee
	database.DB.First(&foundUser, "email = ?", user.Email)
	if foundUser.ID == 0{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Invalid email or password",
		})
		return
	}

	//compare sent in password with saved used password hash
	err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Invalid email or password",
		})
		return
	}
	//tokenString := helpers.GenerateJwtToken(c, foundUser.Name, foundUser.ContactNumber, foundUser.Address, foundUser.Email)
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", foundUser.Token, 3600 * 24, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message":"user has sussessfully loggedIn......!!"})
}
