package jwt

import (
	"time"

	"example.com/config"
	"example.com/internal/user"
	"github.com/golang-jwt/jwt/v5"
)

type jwtRepository struct{}

func NewJWTRepository() Repository {
	return &jwtRepository{}
}

func newToken(u *user.User, duration time.Duration) (string, error) {
	// Generate token claims
	claims := jwt.MapClaims{
		"user_id": u.ID,
		"exp":     time.Now().Add(duration).Unix(),
		"iat":     time.Now().Unix(),
	}

	// Sign token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(config.Environments.JWTSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (repo *jwtRepository) NewTokens(u *user.User) (*JWTTokens, error) {
	// New access token
	accessToken, err := newToken(u, config.Constants.JWTAccessDuration)
	if err != nil {
		return nil, err
	}

	// New refresh token
	refreshToken, err := newToken(u, config.Constants.JWTRefreshDuration)
	if err != nil {
		return nil, err
	}

	return &JWTTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (repo *jwtRepository) VerifyToken(accessToken string) error {
	return nil
}

func (repo *jwtRepository) RefreshTokens(accessToken string) *JWTTokens {
	return nil
}
