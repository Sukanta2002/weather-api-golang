package routes

import (
	"net/http"

	"github.com/Sukanta2002/weather-api-golang/controllers"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

func WeatherRouter(router *mux.Router, rdb *redis.Client) {

	weatherController := &controllers.WeatherController{Rdb: rdb}

	router.HandleFunc("/location/{location}", weatherController.GetWeather).Methods(http.MethodGet)
}
