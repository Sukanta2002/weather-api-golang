package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Sukanta2002/weather-api-golang/utils"
	"github.com/gorilla/mux"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {

	location := mux.Vars(r)["location"]

	apikey := os.Getenv("WEATHER_API_KEY")
	fmt.Println("api key: ", apikey)
	fmt.Println("Location: ", location)

	res, err := http.Get("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/" + location + "?unitGroup=metric&key=" + apikey + "&contentType=json")

	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	var result map[string]interface{}
	data, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(data, &result)

	fmt.Println(res.Status)

	utils.ApiResponce(w, http.StatusOK, result)

}
