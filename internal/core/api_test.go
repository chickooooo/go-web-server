package core

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Sample struct for testing
type JSONResponse struct {
	Message string `json:"message"`
}

// Test happy flow of WriteJSON
func TestWriteJSON(t *testing.T) {
	// Arrange
	rr := httptest.NewRecorder()
	expectedData := JSONResponse{
		Message: "Link created",
	}
	expectedStatus := http.StatusCreated
	expectedContentType := "application/json"

	// Act
	WriteJSON(rr, expectedStatus, expectedData)

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
	var receivedData JSONResponse
	err := json.NewDecoder(rr.Body).Decode(&receivedData)
	if err != nil {
		t.Fatalf("error decoding response body: %v", err)
	}

	if receivedData != expectedData {
		t.Errorf("expected body %+v, got %+v", expectedData, receivedData)
	}
}

// Test error handling of WriteJSON
func TestWriteJSON_ErrorHandling(t *testing.T) {
	// Arrange
	rr := httptest.NewRecorder()
	// Channels can't be JSON encoded; this will trigger an error.
	data := make(chan int)
	expectedStatus := http.StatusInternalServerError
	expectedData := "Something went wrong\n"

	// Act
	WriteJSON(rr, http.StatusOK, data)

	// Assert
	// Check status code
	if rr.Code != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", expectedStatus, rr.Code)
	}

	// Check response body
	body := rr.Body.String()
	if body != expectedData {
		t.Errorf("expected error message '%s', got '%s'", expectedData, body)
	}
}
