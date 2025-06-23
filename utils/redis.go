package utils

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

func InitRedisConnection() *redis.Client {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		panic(fmt.Sprintf("Can't connect to redis %s", err.Error()))
	}

	protocol, err := strconv.Atoi(os.Getenv("REDIS_PROTOCOL"))
	if err != nil {
		panic(fmt.Sprintf("Can't connect to redis %s", err.Error()))
	}

	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
		Protocol: protocol,
	})
}
