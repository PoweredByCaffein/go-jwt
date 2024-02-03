package controller

import (
	"github.com/gin-gonic/gin"
	"go-jwt/pkg/users"
	"net/mail"
)

func ValidateSignUpRequest(c *gin.Context) (requestBody users.SignUpRequest, err error) {
	// Bind the json
	err = c.ShouldBindJSON(&requestBody)

	// Validate email address
	_, err = mail.ParseAddress(requestBody.Email)

	return
}

func ValidateLoginRequest(c *gin.Context) (requestBody users.LoginRequest, err error) {
	// Bind the json
	err = c.ShouldBindJSON(&requestBody)

	// Validate email address
	_, err = mail.ParseAddress(requestBody.Email)

	return
}
