package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test happy flow of HealthHandler
func TestHealthHandler(t *testing.T) {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()

	expectedStatus := http.StatusOK
	expectedContentType := "application/json"
	expectedBody := HealthResponse{
		Status: "Healthy!",
	}

	// Act
	HealthHandler(rr, req)

	// Assert
	// Check status code
	if rr.Code != expectedStatus {
		t.Errorf("expected status %d, got %d", expectedStatus, rr.Code)
	}

	// Check Content-Type header
	contentType := rr.Header().Get("Content-Type")
	if contentType != expectedContentType {
		t.Errorf("expected Content-Type '%s', got '%s'", expectedContentType, contentType)
	}

	// Check JSON response body
	var receivedBody HealthResponse
	err := json.NewDecoder(rr.Body).Decode(&receivedBody)
	if err != nil {
		t.Fatalf("error decoding response body: %v", err)
	}

	if receivedBody != expectedBody {
		t.Errorf("expected body %+v, got %+v", expectedBody, receivedBody)
	}
}
