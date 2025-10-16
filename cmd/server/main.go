package main

import (
	"fmt"
	"net/http"

	"github.com/chickooooo/internal/api"
)

func main() {
	// Create a server multiplexer
	mux := http.NewServeMux()

	// Initialize API routes
	api.InitializeRoutes(mux)

	// Start server
	fmt.Println("HTTP server started at port 8080")
	http.ListenAndServe(":8080", mux)
}
