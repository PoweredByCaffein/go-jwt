package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"go-jwt/pkg/users"
	"net/http"
	"os"
	"strings"
	"time"
)

var MissingAuthTokenMessage = "Authorization token missing!" +
	"Please ensure that you have a valid Authorization token in your request header"

func IsAuthorized(c *gin.Context) {
	// Get the authorization off the request
	tokenString := c.GetHeader("Authorization")
	if strings.TrimSpace(tokenString) == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": MissingAuthTokenMessage})
		c.Abort()
		return
	}

	// Decode and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header)
		}
		return []byte(os.Getenv("APP_KEY")), nil
	})

	if err != nil {
		log.Error().Msgf("Error decoding the token: %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			log.Error().Msgf("Token expired for user [%s]", claims["sub"].(string))
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Validate user
		user, err := users.FetchUserByEmail(claims["sub"].(string))
		if err != nil {
			log.Error().Msgf("Failed to fetch user by email: %s", err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		log.Debug().Msgf("%s logged into the account!", user.FirstName)

		c.Set("user", user)
		c.Next()
	} else {
		log.Error().Msgf("Error during claims")
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
