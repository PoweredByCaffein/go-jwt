package mysql

import (
	"os"

	"github.com/uptrace/bun"
)

func ConnectToUserDatabase(debugMode bool) (*bun.DB, error) {

	svc := Connection{
		Host:     os.Getenv("MYSQL_DATABASE_HOST"),
		Port:     os.Getenv("MYSQL_DATABASE_PORT"),
		Username: os.Getenv("MYSQL_DATABASE_USERNAME"),
		Password: os.Getenv("MYSQL_DATABASE_PASSWORD"),
		Database: os.Getenv("MYSQL_DATABASE_DB"),
		Debug:    debugMode,
	}

	if err := svc.Connect(); err != nil {
		return nil, err
	}

	return svc.Client, nil

}
