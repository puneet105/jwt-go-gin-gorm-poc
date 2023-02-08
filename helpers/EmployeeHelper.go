package helpers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func GeneratePasswordHash(c *gin.Context, password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Failed to hash password",
		})
		return "nil", err
	}
	return string(hash), nil
}
