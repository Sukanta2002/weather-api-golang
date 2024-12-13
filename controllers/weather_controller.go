package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Sukanta2002/weather-api-golang/utils"
	"github.com/gorilla/mux"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {

	location := mux.Vars(r)["location"]

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

	json.Unmarshal(data, &result)

	utils.ApiResponce(w, http.StatusOK, result)

}
