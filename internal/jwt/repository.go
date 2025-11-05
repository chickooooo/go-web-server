package jwt

import "example.com/internal/user"

type Repository interface {
	NewTokens(u *user.User) (*JWTTokens, error)
	VerifyToken(accessToken string) error
	RefreshTokens(accessToken string) *JWTTokens
}
