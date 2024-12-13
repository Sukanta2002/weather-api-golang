package utils

import (
	"encoding/json"
	"net/http"
)

func ApiResponce(w http.ResponseWriter, status int, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	sucess := true
	if status >= 400 {
		sucess = false
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"sucess": sucess,
		"data":   data,
		"status": status,
	})
}
