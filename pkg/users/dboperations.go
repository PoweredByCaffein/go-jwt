package users

import (
	"context"
	"github.com/rs/zerolog/log"
	"go-jwt/pkg/database/mysql"
	models "go-jwt/pkg/models/user"
)

func CreateUser(requestBody SignUpRequest, passwordHash []byte) error {
	// Create connection
	conn, err := mysql.ConnectToDefaultDatabase()
	if err != nil {
		log.Error().Msgf("Failed to connect to DB: %s", err.Error())
		return err
	}

	// Saved data in the database
	_, err = conn.NewInsert().Model(&models.UserModel{
		FirstName: requestBody.FirstName,
		LastName:  requestBody.LastName,
		Email:     requestBody.Email,
		Password:  string(passwordHash),
	}).Exec(context.TODO())

	return err
}
