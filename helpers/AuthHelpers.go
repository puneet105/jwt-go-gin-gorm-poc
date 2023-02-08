package helpers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"time"
)

type SignedDetails struct{
	Name 			string
	ContactNumber	string
	Address			string
	Email			string
	jwt.StandardClaims
}

func GenerateJwtToken(c *gin.Context, name string, contact string, address string, email string) string {
	claims := &SignedDetails{
		Name:          name,
		ContactNumber: contact,
		Address:       address,
		Email:         email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid token",
		})
		return ""
	}
	return tokenString
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string){
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error){
			return []byte(os.Getenv("SECRET")), nil
		},
	)
	if err != nil{
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok{
		msg = fmt.Sprintf("Token Is Invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Unix(){
		msg = fmt.Sprintf("Token Is Expired")
		msg = err.Error()
		return
	}
	fmt.Println("Token Is Valid")
	return claims, msg
}
	/*
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub" : user.ID,
		"exp" : time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	//sign and get complete encoded token
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Invalid token",
		})
		return
	}

	 */

