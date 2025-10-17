package main

import (
	"fmt"
	"net/http"

	"github.com/chickooooo/go-web-server/internal/api"
	"github.com/chickooooo/go-web-server/internal/config"
)

// Start web server
func StartServer(
	configLoader func(string) *config.Config,
	routeInitializer func(*http.ServeMux),
	listenAndServe func(string, http.Handler) error,
) {
	// Load project config
	cfg := configLoader(".env")

	// Create a server multiplexer & initialise routes
	mux := http.NewServeMux()
	routeInitializer(mux)

	// Start the server
	fmt.Printf("HTTP server started at port %s\n", cfg.ServerPort)
	serverAddress := fmt.Sprintf(":%s", cfg.ServerPort)
	listenAndServe(serverAddress, mux)
}

func main() {
	StartServer(
		config.LoadConfig,
		api.InitializeRoutes,
		http.ListenAndServe,
	)
}
