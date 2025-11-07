package utils

import (
	"sync"

	"example.com/internal/handler"
	"example.com/internal/jwt"
	"example.com/internal/user"
)

type Handlers struct {
	Core handler.CoreHandler
	Auth handler.AuthHandler
}

var (
	handlers     *Handlers
	handlersOnce sync.Once
)

// GetHandlers initializes and returns all the handlers
func GetHandlers() *Handlers {
	handlersOnce.Do(func() {
		GolangJWTRepo := jwt.NewJWTRepository()
		userSQLRepo := user.NewSQLRepository()

		jwtService := jwt.NewService(GolangJWTRepo)
		userService := user.NewService(userSQLRepo)

		coreHandler := handler.NewCoreHandler()
		authHandler := handler.NewAuthHandler(jwtService, userService)

		handlers = &Handlers{
			Core: coreHandler,
			Auth: authHandler,
		}
	})
	return handlers
}
