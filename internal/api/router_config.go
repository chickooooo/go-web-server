package api

import "net/http"

// HandlerSet holds all web request handlers
type HandlerSet struct {
	HealthHandler   func(http.ResponseWriter, *http.Request)
	NotFoundHandler func(http.ResponseWriter, *http.Request)
}

// NewEmptyHandlerSet returns an empty handler set.
// Intended use is for mocking handlers in test cases.
func NewEmptyHandlerSet() *HandlerSet {
	return &HandlerSet{}
}

// NewHandlerSet returns a new handler set of actual handlers.
func NewHandlerSet() *HandlerSet {
	return &HandlerSet{
		HealthHandler:   HealthHandler,
		NotFoundHandler: NotFoundHandler,
	}
}
