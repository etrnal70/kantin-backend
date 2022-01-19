package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = id
  // TODO: Is this actually correct ?
	claims["aud"] = "kantin.default_csunj"
	claims["iss"] = "default_csunj"
	claims["iat"] = time.Now().Unix()

	// Set expiry to 12 hour
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authString := strings.Fields(c.GetHeader("Authorization"))
		if len(authString) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header not found",
			})
		}
		token, err := jwt.Parse(authString[1], func(t *jwt.Token) (interface{}, error) {
			// Check signing method
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf(("Invalid Signing Method"))
			}

			// Check validity and expiry
			if _, ok := t.Claims.(jwt.Claims); !ok && !t.Valid {
				return nil, fmt.Errorf(("Token Expired"))
			}

			// TODO: Check other claims (aud, iss, etc)
			secret := []byte(os.Getenv("TOKEN_SECRET"))
			return secret, nil
		})
		if err != nil || !token.Valid {
			fmt.Println(err)
			fmt.Println(token.Valid)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "Session unauthorized",
			})
			return
		}

		c.Next()
	}
}
