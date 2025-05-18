package db

import (
	"github.com/redis/go-redis/v9"
)

// [https://redis.io/docs/latest/develop/clients/go/]
func Conn_Redis() *redis.Client {
	var useDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return useDB
}
