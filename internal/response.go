package internal

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func SendJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
