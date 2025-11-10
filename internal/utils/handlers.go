package utils

import (
	"sync"

	"example.com/internal/handler"
	"example.com/internal/jwt"
	"example.com/internal/product"
	"example.com/internal/user"
)

type Handlers struct {
	Core    handler.CoreHandler
	Auth    handler.AuthHandler
	Product handler.ProductHandler
}

var (
	handlers     *Handlers
	handlersOnce sync.Once
)

// GetHandlers initializes and returns all the handlers
func GetHandlers() *Handlers {
	handlersOnce.Do(func() {
		jwtRepo := jwt.NewJWTRepository()
		userSQLRepo := user.NewSQLRepository()
		productSQLRepo := product.NewSQLRepository()

		jwtService := jwt.NewService(jwtRepo)
		userService := user.NewService(userSQLRepo)
		productService := product.NewService(productSQLRepo)

		coreHandler := handler.NewCoreHandler()
		authHandler := handler.NewAuthHandler(jwtService, userService)
		productHandler := handler.NewProductHandler(productService)

		handlers = &Handlers{
			Core:    coreHandler,
			Auth:    authHandler,
			Product: productHandler,
		}
	})
	return handlers
}
