package user

import (
	"errors"
	"strings"
	"time"
)

type CreateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDTO struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type User struct {
	ID        int
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Validate validates CreateUser model
func (cu *CreateUser) Validate() error {
	// Trim spaces
	cu.Username = strings.TrimSpace(cu.Username)
	cu.Password = strings.TrimSpace(cu.Password)

	// Username and password validation
	if len(cu.Username) < 3 {
		return errors.New("Username should be atleast 3 characters long")
	}
	if len(cu.Password) < 3 {
		return errors.New("Password should be atleast 3 characters long")
	}

	return nil
}
