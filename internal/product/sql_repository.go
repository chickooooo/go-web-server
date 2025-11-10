package product

import "time"

type sqlRepository struct{}

func NewSQLRepository() Repository {
	return &sqlRepository{}
}

func (repo *sqlRepository) Create(cp *CreateProduct) (*Product, error) {
	return &Product{
		ID:        1,
		Name:      "abcd",
		Price:     10,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (repo *sqlRepository) ByID(productId int) (*Product, error) {
	return &Product{
		ID:        1,
		Name:      "abcd",
		Price:     10,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
