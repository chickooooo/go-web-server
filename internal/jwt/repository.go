package jwt

type Repository interface {
	NewTokens(td *TokenData) (*JWTTokens, error)
	VerifyToken(tokenStr string) (*TokenData, error)
}
