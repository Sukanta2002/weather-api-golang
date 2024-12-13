package routes

import "github.com/gorilla/mux"

func SetRoutes() *mux.Router {
	router := mux.NewRouter()

	WeatherRouter(router)

	return router
}
