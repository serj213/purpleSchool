package response

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, data interface{}, statusCode int){
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}