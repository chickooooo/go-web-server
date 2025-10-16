package api

import "net/http"

func InitializeRoutes(mux *http.ServeMux) {
	// Define API routes
	mux.HandleFunc("GET /", healthHandler)
	mux.HandleFunc("GET /health", healthHandler)
}
