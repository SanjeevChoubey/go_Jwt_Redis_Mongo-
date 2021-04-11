package database

import (
	"os"

	"github.com/go-redis/redis/v7"
)

var RedisClient *redis.Client = RedisInstance()

func RedisInstance() *redis.Client {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client := redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}
