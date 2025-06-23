package utils

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

func InitRedisConnection() (*redis.Client, error) {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return nil, err
	}

	protocol, err := strconv.Atoi(os.Getenv("REDIS_PROTOCOL"))
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
		Protocol: protocol,
	})

	return client, nil
}
