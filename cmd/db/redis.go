package db

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(addr string) (*redis.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	redisCocnectionCheck := client.Ping(ctx)
	if redisCocnectionCheck.Err() != nil {
		return nil, redisCocnectionCheck.Err()
	}

	return client, nil

}

func DisconnectRedis(client *redis.Client) error {
	return client.Close()
}
