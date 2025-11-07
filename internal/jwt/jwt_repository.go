package jwt

import (
	"fmt"
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
	TokenType string
	jwt.RegisteredClaims
}

// newToken generates a new token containing tokenData.
// tokenType can be 'access' or 'refresh'
func newToken(td *TokenData, tokenType string) (string, error) {
	// Get token duration from token type
	var duration time.Duration
	switch tokenType {
	case "access":
		duration = config.Constants.JWTAccessDuration
	case "refresh":
		duration = config.Constants.JWTRefreshDuration
	default:
		return "", fmt.Errorf("Invalid token token type '%s'", tokenType)
	}

	// Generate token claims
	now := time.Now()
	claims := Claims{
		TokenData: TokenData{
			UserID: td.UserID,
		},
		TokenType: tokenType,
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
	accessToken, err := newToken(td, "access")
	if err != nil {
		return nil, err
	}

	// New refresh token
	refreshToken, err := newToken(td, "refresh")
	if err != nil {
		return nil, err
	}

	return &JWTTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// VerifyToken verifies the given tokenStr. Returns the encoded token data
// tokenType can be 'access' or 'refresh'
func (repo *jwtRepository) VerifyToken(tokenStr string, tokenType string) (*TokenData, error) {
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
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	// Validate token type
	if claims.TokenType != tokenType {
		return nil, fmt.Errorf("Received '%s' token, expected '%s' token", claims.TokenType, tokenType)
	}

	return &claims.TokenData, nil
}
