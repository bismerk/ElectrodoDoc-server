package database

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var ClientForToken = NewClient(0)

func NewClient(database int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:            "localhost:6379",
		Password:        "",       // no password set
		DB:              database, // use default DB
		MaxRetries:      5,
		MinRetryBackoff: time.Second,
		MaxRetryBackoff: 5 * time.Second,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
}