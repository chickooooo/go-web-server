package main

import (
	"net/http"
	"testing"

	"github.com/chickooooo/go-web-server/internal/api"
	"github.com/chickooooo/go-web-server/internal/config"
)

// TestStartServer verifies that config loading, route initialization, and server start are invoked
func TestStartServer(t *testing.T) {
	// Flags to verify function calls
	configCalled := false
	routesCalled := false
	serverCalled := false

	// Mock config loader
	mockConfigLoader := func(envPath string) *config.Config {
		configCalled = true
		if envPath != ".env" {
			t.Errorf("expected envPath '.env', got '%s'", envPath)
		}
		return &config.Config{
			ServerPort: "1234",
		}
	}

	// Mock route initializer
	mockRouteInitializer := func(mux *http.ServeMux, handlerSet *api.HandlerSet) {
		routesCalled = true
		if mux == nil {
			t.Error("expected non-nil mux")
		}
		if handlerSet == nil {
			t.Error("expected non-nil handlerSet")
		}
	}

	// Mock ListenAndServe
	mockListenAndServe := func(addr string, handler http.Handler) error {
		serverCalled = true
		expectedAddr := ":1234"
		if addr != expectedAddr {
			t.Errorf("expected server address '%s', got '%s'", expectedAddr, addr)
		}
		if handler == nil {
			t.Error("expected non-nil handler")
		}
		return nil // simulate no error
	}

	// Act
	StartServer(mockConfigLoader, mockRouteInitializer, mockListenAndServe)

	// Assert
	if !configCalled {
		t.Error("expected config loader to be called")
	}
	if !routesCalled {
		t.Error("expected route initializer to be called")
	}
	if !serverCalled {
		t.Error("expected server to start")
	}
}
