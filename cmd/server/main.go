package main

import (
	"fmt"
	"net/http"

	"github.com/chickooooo/go-web-server/internal/api"
	"github.com/chickooooo/go-web-server/internal/config"
)

func main() {
	// Load project config
	cfg := config.LoadConfig()

	// Create a server multiplexer
	mux := http.NewServeMux()

	// Initialize API routes
	api.InitializeRoutes(mux)

	// Start the server
	fmt.Printf("HTTP server started at port %s\n", cfg.ServerPort)
	serverAddress := fmt.Sprintf(":%s", cfg.ServerPort)
	http.ListenAndServe(serverAddress, mux)
}
