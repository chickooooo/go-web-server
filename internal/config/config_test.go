package config

import (
	"os"
	"testing"
)

// Test getEnv returns the environment variable when set
func TestGetEnv_ReturnsValue(t *testing.T) {
	// Arrange
	key := "TEST_ENV_VAR"
	expected := "some_value"
	os.Setenv(key, expected)
	defer os.Unsetenv(key)

	// Act
	actual := getEnv(key, "default_value")

	// Assert
	if actual != expected {
		t.Errorf("expected '%s', got '%s'", expected, actual)
	}
}

// Test getEnv returns default when env variable is not set
func TestGetEnv_ReturnsDefault(t *testing.T) {
	// Arrange
	key := "UNSET_ENV_VAR"
	expected := "default_value"
	os.Unsetenv(key)

	// Act
	actual := getEnv(key, expected)

	// Assert
	if actual != expected {
		t.Errorf("expected default value '%s', got '%s'", expected, actual)
	}
}

// Test LoadConfig returns default values when .env file is missing
func TestLoadConfig_DefaultValues(t *testing.T) {
	// Arrange
	os.Unsetenv("SERVER_PORT")
	nonExistentPath := "non-existent-file.env"
	expected := "8080"

	// Act
	cfg := LoadConfig(nonExistentPath)

	// Assert
	if cfg.ServerPort != expected {
		t.Errorf("expected ServerPort '%s', got '%s'", expected, cfg.ServerPort)
	}
}

// Test LoadConfig returns actual values from environment
func TestLoadConfig_ValuesFromEnv(t *testing.T) {
	// Arrange
	expected := "9090"
	os.Setenv("SERVER_PORT", expected)
	defer os.Unsetenv("SERVER_PORT")

	// Act
	cfg := LoadConfig("non-existent.env") // we don't need a real .env file; env var is already set

	// Assert
	if cfg.ServerPort != expected {
		t.Errorf("expected ServerPort '%s', got '%s'", expected, cfg.ServerPort)
	}
}
