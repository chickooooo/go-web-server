package jwt

type Repository interface {
	NewTokens(td *TokenData) (*JWTTokens, error)
	VerifyToken(tokenStr string, tokenType string) (*TokenData, error)
}
