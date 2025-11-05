package auth

import (
	"encoding/json"
	"net/http"

	"example.com/internal/core"
	"example.com/internal/jwt"
	"example.com/internal/user"
)

type Handler interface {
	Register(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	jwtService  jwt.Service
	userService user.Service
}

func NewHandler(jwtService jwt.Service, userService user.Service) Handler {
	return &handler{
		jwtService:  jwtService,
		userService: userService,
	}
}

func (h handler) Register(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var cu user.CreateUser
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&cu)
	if err != nil {
		response := core.ErrorResponse{
			Message: "Cannot parse request body",
			Details: err.Error(),
		}
		core.WriteJSON(w, http.StatusBadRequest, response)
		return
	}

	// Validate request body
	if err = cu.Validate(); err != nil {
		response := core.ErrorResponse{
			Message: "Invalid request body",
			Details: err.Error(),
		}
		core.WriteJSON(w, http.StatusBadRequest, response)
		return
	}

	// Create user
	user, err := h.userService.Create(&cu)
	if err != nil {
		response := core.ErrorResponse{
			Message: "Error creating user",
			Details: err.Error(),
		}
		core.WriteJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Generate JWT tokens
	tokens, err := h.jwtService.NewTokens(user)
	if err != nil {
		response := core.ErrorResponse{
			Message: "Error generating tokens",
			Details: err.Error(),
		}
		core.WriteJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Success response
	core.WriteJSON(w, http.StatusCreated, tokens)
}
