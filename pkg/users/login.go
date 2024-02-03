package users

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"go-jwt/pkg/database/mysql"
	"go-jwt/pkg/helpers"
	models "go-jwt/pkg/models/user"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func Login(request LoginRequest) (statusCode int, response LoginResponse) {
	// Get user details
	user, err := FetchUserByEmail(request.Email)
	if err != nil {
		log.Error().Msgf("Failed to fetch user: %s", err.Error())
	}

	if user == nil {
		log.Warn().Msgf("User with provided email [%s] doesn't exist", user.Email)
		response.Error = fmt.Errorf("invalid username/password, please ensure that you are using the correct combination")
		return http.StatusUnauthorized, response
	}

	// Match password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		log.Warn().Msgf("Incorrect password provided email [%s]", user.Email)
		response.Error = fmt.Errorf("invalid username/password, please ensure that you are using the correct combination")
		return http.StatusUnauthorized, response
	}

	loginDurationStr := os.Getenv("LOGIN_DURATION")
	loginDuration, _ := helpers.ConvertStringToInt(loginDurationStr, 2)

	// Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": request.Email,
		"exp": time.Now().Add(time.Duration(loginDuration) * time.Hour).Unix(),
	})

	appKey := os.Getenv("APP_KEY")
	if appKey == "" {
		response.Error = fmt.Errorf("missing app key")
		return http.StatusInternalServerError, response
	}

	tokenString, err := token.SignedString([]byte(appKey))
	if err != nil {
		log.Error().Msgf("Failed to create signed string: %s", err.Error())
		response.Error = fmt.Errorf("failed to create signed string")
		return http.StatusInternalServerError, response
	}

	response.Success = true
	response.Token = tokenString

	log.Info().Msgf("User logged in successfully!")
	return http.StatusOK, response
}

func FetchUserByEmail(email string) (*models.UserModel, error) {
	db, err := mysql.ConnectToDefaultDatabase()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Warn().Msgf("[FetchUserByEmail] Failed to close database connection: %s", err.Error())
		}
	}()

	var user models.UserModel
	err = db.NewSelect().
		Model(&models.UserModel{}).
		Where("email = ?", email).
		Limit(1).
		Scan(context.TODO(), &user)

	return &user, err
}
