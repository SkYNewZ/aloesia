package controllers

import (
	"encoding/json"
	"net/http"
)

// HealthCheckHandler heath check endpoint
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: implement a better health check
	json.NewEncoder(w).Encode(map[string]string{"ping": "pong"})
}
