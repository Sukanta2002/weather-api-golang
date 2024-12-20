package routes

import (
	"net/http"

	"github.com/Sukanta2002/weather-api-golang/controllers"
	"github.com/gorilla/mux"
)

func WeatherRouter(router *mux.Router) {
	router.HandleFunc("/location/{location}", controllers.GetWeather).Methods(http.MethodGet)
}
