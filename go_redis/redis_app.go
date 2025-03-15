package main

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// ServerClient function  
func ServerClient() {
	up_token := os.Getenv("UP_API_KEY")
	server_name := os.Getenv("SERVER_NAME")

	rdb := redis.NewClient(&redis.Options{
		Addr:     server_name + ":6379",
		Password: "",
		DB:       0,
	})

	err := rdb.Set(ctx, "up_token", up_token, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "up_token").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

}

// main function  
func main() {
	fmt.Println("Testing Redis app...")

	ServerClient()
}
