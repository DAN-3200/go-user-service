package db

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// [https://redis.io/docs/latest/develop/clients/go/]
func Conn_Redis() *redis.Client {
	conn := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if err := conn.Ping(context.Background()).Err(); err != nil {
		fmt.Printf("Erro: %v", err)
		return nil
	}

	return conn
}
