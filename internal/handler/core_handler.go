package handler

import (
	"net/http"

	"example.com/internal/core"
)

type CoreHandler interface {
	Health(w http.ResponseWriter, r *http.Request)
}

type coreHandler struct{}

func NewCoreHandler() CoreHandler {
	return &coreHandler{}
}

// Health checks if the server is healthy
func (c coreHandler) Health(w http.ResponseWriter, r *http.Request) {
	healthy := struct {
		Status string `json:"status"`
	}{"Healthy!"}
	core.WriteJSON(w, http.StatusCreated, healthy)
}
