package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Create a response
	response := HealthResponse{Status: "Healthy!"}

	// Encode the respose to JSON and write to response writer
	if err := json.NewEncoder(w).Encode(response); err != nil {
		// Log error
		fmt.Printf("Error while encoding to JSON. Data: %+v. Error: %v.\n", response, err)

		// Send error response
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
}
