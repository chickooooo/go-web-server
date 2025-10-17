package api

import "net/http"

func InitializeRoutes(rootMux *http.ServeMux) {
	// Subrouter
	apiMux := http.NewServeMux()

	// Define API routes
	apiMux.HandleFunc("GET /health", HealthHandler)

	// Add "/api/v1" prefix for each route
	stripHandler := http.StripPrefix("/api/v1", apiMux)
	rootMux.Handle("/api/v1/", stripHandler)
}
