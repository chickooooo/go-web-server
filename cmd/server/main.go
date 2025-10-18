package main

import (
	"fmt"
	"net/http"

	"github.com/chickooooo/go-web-server/internal/api"
	"github.com/chickooooo/go-web-server/internal/config"
)

// Start web server
func StartServer(
	envConfigLoader func(string) *config.EnvConfig,
	routeInitializer func(*http.ServeMux, *api.HandlerSet),
	listenAndServe func(string, http.Handler) error,
) {
	// Load project config
	envConfig := envConfigLoader(".env")

	// Create a server multiplexer & initialise routes
	mux := http.NewServeMux()
	handlerSet := api.NewHandlerSet()
	routeInitializer(mux, handlerSet)

	// Start the server
	fmt.Printf("HTTP server started at port %s\n", envConfig.ServerPort)
	serverAddress := fmt.Sprintf(":%s", envConfig.ServerPort)
	listenAndServe(serverAddress, mux)
}

func main() {
	StartServer(
		config.LoadEnvConfig,
		api.InitializeRoutes,
		http.ListenAndServe,
	)
}
