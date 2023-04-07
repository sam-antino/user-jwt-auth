package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"user-jwt-auth/repository"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ReqAuthorization(c *gin.Context) {
	t := c.Request.Header["Authorization"]

	// if token is not passed in the header
	if len(t) == 0 {
		AbortMiddleware(c)
	}

	tokenString := t[0]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		AbortMiddleware(c)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check expiry
		if claims["exp"].(float64) <= float64(time.Now().Unix()) {
			AbortMiddleware(c)
		}

		// check email
		_, err = repository.GetUser(claims["email"].(string))
		if err != nil {
			AbortMiddleware(c)
		}

		// pass email in the header for further services
		c.Request.Header.Add("user-email", claims["email"].(string))
		c.Next()

	} else {
		AbortMiddleware(c)
	}

}

func AbortMiddleware(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error":   "something went wrong",
		"message": "Authentication failed",
	})
	c.Abort()
}
