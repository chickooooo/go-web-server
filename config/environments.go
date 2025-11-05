package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type environments struct {
	JWTSecret []byte
}

var (
	Environments     environments
	EnvironmentsOnce sync.Once
)

// LoadEnv loads environment variables
func LoadEnv(envPath string) {
	// Initialize only once
	EnvironmentsOnce.Do(func() {
		// Load .env file
		err := godotenv.Load(envPath)
		if err != nil {
			panic(fmt.Sprintf("Cannot load environment variables: %v", err))
		}

		// Setup environment variables
		Environments = environments{
			JWTSecret: []byte(getEnv("JWT_SECRET")),
		}
	})
}

func getEnv(key string) string {
	// Try to get the key from environment
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	// Panic if key not present
	panic(fmt.Sprintf("Missing key %q from environment\n", key))
}
