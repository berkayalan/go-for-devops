package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	redis "github.com/redis/go-redis/v9"
)

type User struct {
	ID         string
	Name       string "json:name"
	Age        int    "json:age"
	Occupation string "json:occupation"
}

func main() {

	redis_client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ping, err := redis_client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(ping)

	userID := uuid.NewString()
	jsonString, err := json.Marshal(User{
		ID:         userID,
		Name:       "Berkay",
		Age:        28,
		Occupation: "Cloud Engineer",
	})
	if err != nil {
		fmt.Printf("Failed to marshall : %s", err)
		return
	}

	err = redis_client.Set(context.Background(), string(userID), jsonString, 0).Err()
	if err != nil {
		fmt.Printf("Failed to set value : %s", err)
		return
	}

	val, err := redis_client.Get(context.Background(), string(userID)).Result()
	if err != nil {
		fmt.Printf("Failed to get value : %s", err)
		return
	}

	fmt.Printf("Value retrieved from Redis: %s\n", val)

}
