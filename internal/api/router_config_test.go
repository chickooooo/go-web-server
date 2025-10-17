package api

import (
	"reflect"
	"testing"
)

func TestNewEmptyHandlerSet(t *testing.T) {
	// Act
	handlerSet := NewEmptyHandlerSet()

	// Assert
	if handlerSet == nil {
		t.Fatal("expected non-nil HandlerSet, got nil")
	}
	if handlerSet.HealthHandler != nil {
		t.Errorf("expected nil HealthHandler, got non-nil")
	}
	if handlerSet.NotFoundHandler != nil {
		t.Errorf("expected nil NotFoundHandler, got non-nil")
	}
}

func TestNewHandlerSet(t *testing.T) {
	// Act
	handlerSet := NewHandlerSet()

	// Assert
	if handlerSet == nil {
		t.Fatal("expected non-nil HandlerSet, got nil")
	}
	if handlerSet.HealthHandler == nil {
		t.Error("expected non-nil HealthHandler, got nil")
	}
	if handlerSet.NotFoundHandler == nil {
		t.Error("expected non-nil NotFoundHandler, got nil")
	}

	// Verify correct mapping
	if reflect.ValueOf(handlerSet.HealthHandler).Pointer() != reflect.ValueOf(HealthHandler).Pointer() {
		t.Errorf("HealthHandler does not match expected function")
	}

	if reflect.ValueOf(handlerSet.NotFoundHandler).Pointer() != reflect.ValueOf(NotFoundHandler).Pointer() {
		t.Errorf("NotFoundHandler does not match expected function")
	}
}
