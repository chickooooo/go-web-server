package jwt

import "example.com/internal/user"

type Service interface {
	NewTokens(u *user.User) (*JWTTokens, error)
	VerifyToken(accessToken string) error
	RefreshTokens(accessToken string) *JWTTokens
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) NewTokens(u *user.User) (*JWTTokens, error) {
	return s.repo.NewTokens(u)
}

func (s *service) VerifyToken(accessToken string) error {
	return s.repo.VerifyToken(accessToken)
}

func (s *service) RefreshTokens(accessToken string) *JWTTokens {
	return s.RefreshTokens(accessToken)
}
