package redis

import (
	"go-jwt/pkg/helpers"
	"os"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func ConnectToDefaultRedisDB() (*redis.Client, error) {

	db, err := helpers.ConvertStringToInt(os.Getenv("REDIS_DB"), 0)
	if err != nil {
		log.Warn().Msgf("Failed to convert [%s] to int: %s", os.Getenv("REDIS_DB"), err.Error())
	}

	svc := Connection{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Database: db,
	}

	if err = svc.Connect(); err != nil {
		return nil, err
	}

	return svc.Client, nil
}
