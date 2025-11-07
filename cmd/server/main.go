package main

import (
	"fmt"
	"net/http"

	"example.com/config"
	"example.com/internal/apis"
)

func main() {
	// Load environment
	config.LoadEnv(".env.example")

	// Initialize routes
	mux := apis.InitializeRoutes()

	// Start server
	fmt.Println("Server started at port 8080")
	http.ListenAndServe(":8080", mux)
}
