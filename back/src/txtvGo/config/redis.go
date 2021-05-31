package config

import (
	"fmt"
	"github.com/go-redis/redis"
)

func init () {
	fmt.Println("----------------> redis.go read successful")
	connectRedis()
}

func connectRedis () {
	rds := redis.NewClient(&redis.Options{
		Addr:     "192.168.149.128:6379",
	})
	_, err := rds.Ping().Result()
	fmt.Println(err)
}
