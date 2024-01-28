package redis

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func (s *Connection) Connect() error {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", s.Host, s.Port),
		Password: s.Password,
		DB:       s.Database,
	})

	s.Client = client

	return nil

}
