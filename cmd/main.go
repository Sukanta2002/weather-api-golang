package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Sukanta2002/weather-api-golang/routes"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	if redisHost == "" || redisPort == "" {
		redisHost = "localhost" // Fallback to localhost for local testing
		redisPort = "6379"
	}
	add := redisHost + ":" + redisPort

	rdb := redis.NewClient(&redis.Options{
		Addr:     add,
		Password: "",
		DB:       0,
	})

	// Test the connection
	ctx := context.Background()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	Port := os.Getenv("PORT")

	router := routes.SetRoutes(rdb)

	fmt.Println("Starting the server at: ", Port)
	http.ListenAndServe(":"+Port, router)
}
