package apis

import (
	"net/http"

	"example.com/internal/utils"
)

func InitializeRoutes() *http.ServeMux {
	// Initialize handlers
	handlers := utils.GetHandlers()

	// Define routes
	mux := http.NewServeMux()
	// auth
	mux.HandleFunc("POST /api/v1/auth/register", handlers.Auth.Register)
	mux.HandleFunc("POST /api/v1/auth/tokens/refresh", handlers.Auth.RefreshTokens)
	// product
	mux.HandleFunc("POST /api/v1/products", handlers.Product.Create)
	mux.HandleFunc("GET /api/v1/products/{productId}", handlers.Product.ById)
	// core
	mux.HandleFunc("/", handlers.Core.Health)

	return mux
}
