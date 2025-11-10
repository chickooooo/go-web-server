package product

import (
	"errors"
	"strings"
	"time"
)

// CreateProduct model is used to create a new product
type CreateProduct struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// ProductDTO model is a subset of Product model.
// It defines the public fields of Product model
type ProductDTO struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// Core Product model
type Product struct {
	ID        int
	Name      string
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Validate validates CreateProduct model
func (cp *CreateProduct) Validate() error {
	// Trim spaces
	cp.Name = strings.TrimSpace(cp.Name)

	// Validations
	if len(cp.Name) < 3 {
		return errors.New("Name should be atleast 3 characters long")
	}
	if cp.Price <= 0 {
		return errors.New("Price should be more than 0")
	}

	return nil
}
