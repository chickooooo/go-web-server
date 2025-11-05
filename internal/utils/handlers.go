package utils

import (
	"sync"

	"example.com/internal/auth"
	"example.com/internal/jwt"
	"example.com/internal/user"
)

type Handlers struct {
	Auth auth.Handler
}

var (
	handlers     *Handlers
	handlersOnce sync.Once
)

func GetHandlers() *Handlers {
	handlersOnce.Do(func() {
		GolangJWTRepo := jwt.NewJWTRepository()
		userSQLRepo := user.NewSQLRepository()

		jwtService := jwt.NewService(GolangJWTRepo)
		userService := user.NewService(userSQLRepo)

		authHandler := auth.NewHandler(jwtService, userService)

		handlers = &Handlers{
			Auth: authHandler,
		}
	})
	return handlers
}
