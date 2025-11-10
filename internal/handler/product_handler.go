package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/internal/core"
	"example.com/internal/product"
)

type ProductHandler interface {
	// Create creates a product and returns it
	Create(w http.ResponseWriter, r *http.Request)

	// ById gets a product using productID and returns it
	// Returns 204 No Content is no matching product found
	ById(w http.ResponseWriter, r *http.Request)
}

type productHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) ProductHandler {
	return &productHandler{
		productService: productService,
	}
}

func (h productHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var cp product.CreateProduct
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&cp)
	if err != nil {
		response := core.ErrorResponse{
			Message: "Cannot parse request body",
			Details: err.Error(),
		}
		core.WriteJSON(w, http.StatusBadRequest, response)
		return
	}

	// Validate request body
	if err = cp.Validate(); err != nil {
		response := core.ErrorResponse{
			Message: "Invalid request body",
			Details: err.Error(),
		}
		core.WriteJSON(w, http.StatusBadRequest, response)
		return
	}

	// Create product
	product, err := h.productService.Create(&cp)
	if err != nil {
		response := core.ErrorResponse{
			Message: "Error creating product",
			Details: err.Error(),
		}
		core.WriteJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Success response
	productDTO := h.productService.ProductToDTO(product)
	core.WriteJSON(w, http.StatusCreated, productDTO)
}

func (h productHandler) ById(w http.ResponseWriter, r *http.Request) {
	// Parse path parameter "productId"
	productIdStr := r.PathValue("productId")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		response := core.ErrorResponse{
			Message: "Invalid productId",
			Details: err.Error(),
		}
		core.WriteJSON(w, http.StatusBadRequest, response)
		return
	}

	// Get product using ID
	product, err := h.productService.ByID(productId)
	if err != nil {
		response := core.ErrorResponse{
			Message: "Error getting product",
			Details: err.Error(),
		}
		core.WriteJSON(w, http.StatusInternalServerError, response)
		return
	} else if product == nil { // No matching product
		response := struct{}{}
		core.WriteJSON(w, http.StatusNoContent, response)
		return
	}

	// Success response
	productDTO := h.productService.ProductToDTO(product)
	core.WriteJSON(w, http.StatusOK, productDTO)
}
