package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data any) {
	// Encode the response to JSON into a buffer first
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		// Log error
		fmt.Printf("Error while encoding to JSON. Data: %+v. Error: %v.\n", data, err)

		// Set error headers and send generic error
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	// Set content type and write status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonBytes)
}
