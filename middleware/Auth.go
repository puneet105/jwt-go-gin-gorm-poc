package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/puneet105/jwt-go-gin-gorm-poc/helpers"
	"net/http"
)

func RequireAuth(c *gin.Context){
	fmt.Println("In Middleware...!!")
	clientToken := c.Request.Header.Get("token")
	if clientToken == ""{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("No Authorization Header Provided...!!"),
		})
		c.Abort()
		return
	}

	claims, err := helpers.ValidateToken(clientToken)
	if err != ""{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()
		return
	}
	c.Set("name", claims.Name)
	c.Set("contact_number", claims.ContactNumber)
	c.Set("address",claims.Address)
	c.Set("email", claims.Email)
	c.Next()
}

/*
//get the cookie off request
	tokenString, err := c.Cookie("Authorization")
	if err != nil{
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//check the expiration of cookie
		if float64(time.Now().Unix()) > claims["exp"].(float64){
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//find user with token sub
		var user models.User
		database.DB.First(&user, claims["sub"])
		if user.ID == 0{
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//attach to request
		c.Set("user", user)
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
 */