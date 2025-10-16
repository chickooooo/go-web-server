package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config contains project configurations
type Config struct {
	ServerPort string
}

// LoadConfig loads environment variables and returns a Config struct
func LoadConfig() *Config {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	return &Config{
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
