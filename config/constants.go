package config

import "time"

type constants struct {
	JWTAccessDuration  time.Duration
	JWTRefreshDuration time.Duration
}

var Constants = constants{
	// JWT
	JWTAccessDuration:  time.Duration(15 * time.Minute),   // 15 minutes
	JWTRefreshDuration: time.Duration(7 * 24 * time.Hour), // 7 days
}
