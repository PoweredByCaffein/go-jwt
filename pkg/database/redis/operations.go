package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func SetValue(client *redis.Client, key string, value interface{}) error {
	ctx := context.Background()
	if err := client.Set(ctx, key, value, 0).Err(); err != nil {
		return err
	}

	return nil
}

func DeleteValue(client *redis.Client, key string) error {
	ctx := context.Background()
	if err := client.Del(ctx, key).Err(); err != nil {
		return err
	}

	return nil
}

func ReadValue(client *redis.Client, key string) (interface{}, error) {
	ctx := context.Background()

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return val, nil
}
