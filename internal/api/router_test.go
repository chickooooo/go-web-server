package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockHandlerSet returns a set of mock handlers
func MockHandlerSet() *HandlerSet {
	handlerSet := NewEmptyHandlerSet()

	// Set mock handlers
	// Each mock handler returns the corresponding handler name
	handlerSet.HealthHandler = func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HealthHandler"))
	}
	handlerSet.NotFoundHandler = func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("NotFoundHandler"))
	}
	//? define new mock handlers here

	return handlerSet
}

// TestInitializeRoutes verifies that each defined route calls the correct handler.
func TestInitializeRoutes(t *testing.T) {
	// Arrange
	rootMux := http.NewServeMux()
	handlerSet := MockHandlerSet()
	InitializeRoutes(rootMux, handlerSet)

	tests := []struct {
		name         string
		method       string
		target       string
		expectedBody string
	}{
		{
			name:         "GET /health is handled by HealthHandler",
			method:       http.MethodGet,
			target:       "/api/v1/health",
			expectedBody: "HealthHandler",
		},
		{
			name:         "GET /abcd is handled by NotFoundHandler",
			method:       http.MethodGet,
			target:       "/api/v1/abcd",
			expectedBody: "NotFoundHandler",
		},
		//? add new test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			req := httptest.NewRequest(tt.method, tt.target, nil)
			rr := httptest.NewRecorder()

			// Act
			rootMux.ServeHTTP(rr, req)
			body := rr.Body.String()

			// Assert
			if body != tt.expectedBody {
				t.Errorf("expected handler %s, got %s", tt.expectedBody, body)
			}
		})
	}
}
