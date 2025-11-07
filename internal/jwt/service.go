package jwt

type Service interface {
	NewTokens(td *TokenData) (*JWTTokens, error)
	VerifyToken(tokenStr string) (*TokenData, error)
	RefreshTokens(refreshToken string) (*JWTTokens, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// NewTokens generate and returns new JWT token pair
func (s *service) NewTokens(td *TokenData) (*JWTTokens, error) {
	return s.repo.NewTokens(td)
}

// VerifyToken verifies the given tokenStr. Returns the encoded token data
func (s *service) VerifyToken(tokenStr string) (*TokenData, error) {
	return s.repo.VerifyToken(tokenStr)
}

// RefreshTokens validates the given refreshToken.
// If the token is valid, it returns a new JWT token pair
func (s *service) RefreshTokens(refreshToken string) (*JWTTokens, error) {
	// Verify refresh token
	tokenData, err := s.repo.VerifyToken(refreshToken)
	if err != nil {
		return nil, err
	}

	// Generate and return new access & refreh tokens
	return s.repo.NewTokens(tokenData)
}
