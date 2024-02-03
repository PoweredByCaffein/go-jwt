package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go-jwt/pkg/constants"
	"go-jwt/pkg/users"
	"net/http"
)

func SignUp(c *gin.Context) {

	// Get details from the request body
	requestBody, err := ValidateSignUpRequest(c)
	if err != nil {
		log.Error().Msgf("Failed to bind request: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": constants.StatusMessageMap[http.StatusBadRequest],
		})
		return
	}

	// Create users
	status, err := users.SignUp(requestBody)
	if err != nil {
		c.JSON(status, gin.H{
			"error": constants.StatusMessageMap[status],
		})
		return
	}

	c.JSON(status, gin.H{
		"success": true,
		"message": "User added successfully",
	})

}
