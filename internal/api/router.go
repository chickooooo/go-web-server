package api

import "net/http"

func InitializeRoutes(rootMux *http.ServeMux, handlerSet *HandlerSet) {
	// Subrouter
	apiMux := http.NewServeMux()

	// Define API routes
	apiMux.HandleFunc("GET /health", handlerSet.HealthHandler)

	// Catch-all unmatched API routes
	apiMux.HandleFunc("/", handlerSet.NotFoundHandler)

	// Add "/api/v1" prefix for each route
	stripHandler := http.StripPrefix("/api/v1", apiMux)
	rootMux.Handle("/api/v1/", stripHandler)
}
