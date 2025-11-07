package handler

import (
	"encoding/json"
	"net/http"

	"example.com/internal/core"
	"example.com/internal/jwt"
	"example.com/internal/user"
)

type AuthHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	RefreshTokens(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	jwtService  jwt.Service
	userService user.Service
}

func NewAuthHandler(jwtService jwt.Service, userService user.Service) AuthHandler {
	return &authHandler{
		jwtService:  jwtService,
		userService: userService,
	}
}

// Register creates a new user, generate JWT tokens for that user and returns them
func (h authHandler) Register(w http.ResponseWriter, r *http.Request) {
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
	tokenData := jwt.TokenData{
		UserID: user.ID,
	}
	tokens, err := h.jwtService.NewTokens(&tokenData)
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

// RefreshTokens refresh JWT tokens for that user and returns them
func (h authHandler) RefreshTokens(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var requestBody struct {
		RefreshToken string `json:"refresh_token"`
	}
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		response := core.ErrorResponse{
			Message: "Cannot parse request body",
			Details: err.Error(),
		}
		core.WriteJSON(w, http.StatusBadRequest, response)
		return
	}

	// Refresh JWT tokens
	tokens, err := h.jwtService.RefreshTokens(requestBody.RefreshToken)
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
