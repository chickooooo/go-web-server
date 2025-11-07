package jwt

import (
	"time"

	"example.com/config"
	"github.com/golang-jwt/jwt/v5"
)

type jwtRepository struct{}

// NewJWTRepository create and returns new JWT Repository
func NewJWTRepository() Repository {
	return &jwtRepository{}
}

// Claims defines the structure for JWT claims
type Claims struct {
	TokenData
	jwt.RegisteredClaims
}

// newToken generates new token that will contain tokenData.
// duration define the TTL of the token.
func newToken(td *TokenData, duration time.Duration) (string, error) {
	// Generate token claims
	now := time.Now()
	claims := Claims{
		TokenData: TokenData{
			UserID: td.UserID,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(duration)),
		},
	}

	// Sign token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(config.Environments.JWTSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// NewTokens generate and returns new JWT token pair
func (repo *jwtRepository) NewTokens(td *TokenData) (*JWTTokens, error) {
	// New access token
	accessToken, err := newToken(td, config.Constants.JWTAccessDuration)
	if err != nil {
		return nil, err
	}

	// New refresh token
	refreshToken, err := newToken(td, config.Constants.JWTRefreshDuration)
	if err != nil {
		return nil, err
	}

	return &JWTTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// VerifyToken verifies the given tokenStr. Returns the encoded token data
func (repo *jwtRepository) VerifyToken(tokenStr string) (*TokenData, error) {
	// Parse the token with a key lookup function
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&Claims{},
		func(token *jwt.Token) (any, error) {
			// Ensure the signing method is HMAC and expected
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenInvalidClaims
			}
			return config.Environments.JWTSecret, nil
		})
	if err != nil {
		return nil, err
	}

	// Extract and validate claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return &claims.TokenData, nil
	}

	return nil, jwt.ErrTokenInvalidClaims
}
