package mysql

import (
	"os"

	"github.com/uptrace/bun"
)

var (
	DbDebugMode  bool
	SetDebugMode int
)

func ConnectToDefaultDatabase() (*bun.DB, error) {

	if SetDebugMode == 0 {
		if dbDebugMode := os.Getenv("DB_DEBUG_MODE"); dbDebugMode == "true" {
			DbDebugMode = true
		}
	}

	svc := Connection{
		Host:     os.Getenv("MYSQL_DATABASE_HOST"),
		Port:     os.Getenv("MYSQL_DATABASE_PORT"),
		Username: os.Getenv("MYSQL_DATABASE_USERNAME"),
		Password: os.Getenv("MYSQL_DATABASE_PASSWORD"),
		Database: os.Getenv("MYSQL_DATABASE_DB"),
		Debug:    DbDebugMode,
	}

	if err := svc.Connect(); err != nil {
		return nil, err
	}

	return svc.Client, nil

}
