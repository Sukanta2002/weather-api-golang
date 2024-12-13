package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Sukanta2002/weather-api-golang/routes"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Port := os.Getenv("PORT")

	router := routes.SetRoutes()

	fmt.Println("Starting the server at: ", Port)
	http.ListenAndServe(":"+Port, router)
}
