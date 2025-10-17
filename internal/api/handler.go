package api

import (
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type HealthResponse struct {
	Status string `json:"status"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	// Healthy response
	response := HealthResponse{Status: "Healthy"}
	WriteJSON(w, http.StatusOK, response)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	// Not found response
	response := ErrorResponse{Message: "Not found"}
	WriteJSON(w, http.StatusNotFound, response)
}
