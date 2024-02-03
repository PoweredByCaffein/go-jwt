package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go-jwt/pkg/constants"
	"go-jwt/pkg/helpers"
	"go-jwt/pkg/users"
	"net/http"
	"os"
)

var LoginDuration int

func Login(c *gin.Context) {
	// Validate request
	requestBody, err := ValidateLoginRequest(c)
	if err != nil {
		log.Error().Msgf("Failed to bind request: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": constants.StatusMessageMap[http.StatusBadRequest],
		})
		return
	}

	loginDurationStr := os.Getenv("LOGIN_DURATION")
	LoginDuration, _ = helpers.ConvertStringToInt(loginDurationStr, 2)

	// Login
	status, response := users.Login(requestBody)

	// Return the token in cookie
	//if status == http.StatusOK {
	//	c.SetSameSite(http.SameSiteLaxMode)
	//	c.SetCookie("Authorization", response.Token, LoginDuration, "", "", false, true)
	//}

	c.JSON(status, response)

}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "I am logged in!",
	})
}
