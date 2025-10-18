package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// EnvConfig contains project environment configurations
type EnvConfig struct {
	ServerPort string
}

// LoadEnvConfig loads environment variables and returns a EnvConfig struct
func LoadEnvConfig(envPath string) *EnvConfig {
	// Load environment variables
	if err := godotenv.Load(envPath); err != nil {
		log.Println("No .env file found, using default values")
	}

	return &EnvConfig{
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
}

// getEnv tries to get the value of key from the environment.
// If the key is not present, defaultVal is returned
func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
