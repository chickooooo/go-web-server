package product

type Repository interface {
	// Create creates a new product record in the database
	Create(cp *CreateProduct) (*Product, error)

	// ByID gets a product using the given productId
	// Returns nil if no matching product found
	ByID(productId int) (*Product, error)
}
