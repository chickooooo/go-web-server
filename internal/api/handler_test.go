package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func verifyStatusCode(t *testing.T, received, expected int) {
	if received != expected {
		t.Errorf("expected status %d, got %d", expected, received)
	}
}

func verifyContentType(t *testing.T, received, expected string) {
	if received != expected {
		t.Errorf("expected Content-Type '%s', got '%s'", expected, received)
	}
}

// Test happy flow of HealthHandler
func TestHealthHandler(t *testing.T) {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()

	expectedStatus := http.StatusOK
	expectedContentType := "application/json"
	expectedBody := HealthResponse{
		Status: "Healthy",
	}

	// Act
	HealthHandler(rr, req)

	// Assert
	// Check status code
	verifyStatusCode(t, rr.Code, expectedStatus)

	// Check Content-Type header
	contentType := rr.Header().Get("Content-Type")
	verifyContentType(t, contentType, expectedContentType)

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

// Test happy flow of NotFoundHandler
func TestNotFoundHandler(t *testing.T) {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/kkkk", nil)
	rr := httptest.NewRecorder()

	expectedStatus := http.StatusNotFound
	expectedContentType := "application/json"
	expectedBody := ErrorResponse{
		Message: "Not found",
	}

	// Act
	NotFoundHandler(rr, req)

	// Assert
	// Check status code
	verifyStatusCode(t, rr.Code, expectedStatus)

	// Check Content-Type header
	contentType := rr.Header().Get("Content-Type")
	verifyContentType(t, contentType, expectedContentType)

	// Check JSON response body
	var receivedBody ErrorResponse
	err := json.NewDecoder(rr.Body).Decode(&receivedBody)
	if err != nil {
		t.Fatalf("error decoding response body: %v", err)
	}

	if receivedBody != expectedBody {
		t.Errorf("expected body %+v, got %+v", expectedBody, receivedBody)
	}
}
