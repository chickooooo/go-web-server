package product

type Service interface {
	// Create creates a new product
	Create(cp *CreateProduct) (*Product, error)

	// ByID gets a product using the given productId
	// Returns nil if no matching product found
	ByID(productId int) (*Product, error)

	// ProductToDTO converts Product model to ProductDTO model
	ProductToDTO(p *Product) ProductDTO
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(cp *CreateProduct) (*Product, error) {
	return s.repo.Create(cp)
}

func (s *service) ByID(productId int) (*Product, error) {
	return s.repo.ByID(productId)
}

func (s *service) ProductToDTO(p *Product) ProductDTO {
	return ProductDTO{
		ID:    p.ID,
		Name:  p.Name,
		Price: p.Price,
	}
}
