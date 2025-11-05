package main

import (
	"fmt"
	"net/http"

	"example.com/config"
	"example.com/internal/core"
	"example.com/internal/utils"
)

// Health is an http handler used to check if the server is helthy
func Health(w http.ResponseWriter, r *http.Request) {
	healthy := struct {
		Status string `json:"status"`
	}{"Healthy!"}
	core.WriteJSON(w, http.StatusCreated, healthy)
}

func main() {
	// Load environment
	config.LoadEnv(".env.example")
	// Initialize handlers
	handlers := utils.GetHandlers()

	// Define routes
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/register", handlers.Auth.Register)
	mux.HandleFunc("/", Health)

	// Start server
	fmt.Println("Server started at port 8080")
	http.ListenAndServe(":8080", mux)
}
