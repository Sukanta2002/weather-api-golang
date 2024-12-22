package main

import (
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

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	Port := os.Getenv("PORT")

	router := routes.SetRoutes(rdb)

	fmt.Println("Starting the server at: ", Port)
	http.ListenAndServe(":"+Port, router)
}
