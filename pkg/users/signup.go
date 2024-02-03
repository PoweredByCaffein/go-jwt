package users

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// SignUp takes the gin.Context and creates the users account
func SignUp(requestBody SignUpRequest) (int, error) {

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), 10)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to generate password hash: %s", err.Error())
	}

	if err = CreateUser(requestBody, hash); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
