package redis

import (
	"github.com/redis/go-redis/v9"
)

type Connection struct {
	Host     string
	Port     string
	Username string
	Password string
	Database int
	Debug    bool

	Client *redis.Client
}
