package config

import (
	"fmt"
	"github.com/go-redis/redis"
)

func init () {
	fmt.Println("----------------> redis.go read successful")
	connectRedis()
}

var Redis *redis.Client

func connectRedis () {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "192.168.149.128:6379",
	})
	pong, err := Redis.Ping().Result()
	fmt.Println(pong, err)
}
