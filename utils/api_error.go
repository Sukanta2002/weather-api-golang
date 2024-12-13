package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ApiError(w http.ResponseWriter, status int, message string) {

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	if status > 499 {
		log.Println(message)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"sucess":  false,
		"data":    nil,
		"status":  status,
		"message": message,
	})

}
