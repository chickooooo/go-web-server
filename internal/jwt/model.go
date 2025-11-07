package jwt

// JWT tokens pair (access & refresh)
type JWTTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Defines the data that will be stored in JWT token
type TokenData struct {
	UserID int `json:"user_id"`
}
