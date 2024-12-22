package routes

import (
	"net/http"

	"github.com/Sukanta2002/weather-api-golang/utils"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

func SetRoutes(rdb *redis.Client) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.ApiResponce(w, http.StatusOK, map[string]interface{}{
			"message": "Weather API in Go",
		})
	})

	WeatherRouter(router, rdb)

	return router
}
