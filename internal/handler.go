package internal

import "net/http"

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Service is healthy",
		Status:  "ok",
	}
	SendJSON(w, http.StatusOK, response)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Hello, World!",
		Status:  "success",
	}
	SendJSON(w, http.StatusOK, response)
}
