package controllers

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Sukanta2002/weather-api-golang/utils"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

type WeatherController struct {
	Rdb *redis.Client
}

func (wc *WeatherController) GetWeather(w http.ResponseWriter, r *http.Request) {

	location := mux.Vars(r)["location"]

	if location == "" {
		utils.ApiError(w, http.StatusBadRequest, "Enter a location")
		return
	}

	ctx := context.Background()

	val, err := wc.Rdb.Get(ctx, location).Result()

	if err == nil {

		var cachedData map[string]interface{}
		err = json.Unmarshal([]byte(val), &cachedData)
		if err != nil {
			log.Println(err.Error())
			utils.ApiError(w, http.StatusInternalServerError, "Failed to parse cached data")
			return
		}
		log.Println("sending the casheed data")
		utils.ApiResponce(w, http.StatusOK, cachedData)
		return
	}

	apikey := os.Getenv("WEATHER_API_KEY")

	res, err := http.Get("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/" + location + "?unitGroup=metric&key=" + apikey + "&contentType=json")

	if err != nil {
		log.Println(err.Error())
		utils.ApiError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
		log.Println(err.Error())
		utils.ApiError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if res.StatusCode != 200 {
		log.Println("Status code: ", res.StatusCode)
		log.Println("res: ", string(data))
		utils.ApiError(w, res.StatusCode, string(data))
		return
	}

	var result map[string]interface{}

	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Println(err.Error())
		utils.ApiError(w, http.StatusInternalServerError, "Failed to marshal data")
		return
	}

	err = wc.Rdb.Set(ctx, location, data, time.Minute*1).Err()
	if err != nil {
		log.Println(err.Error())
		utils.ApiError(w, http.StatusInternalServerError, "Failed to cache weather data")
		return
	}

	utils.ApiResponce(w, http.StatusOK, result)

}
