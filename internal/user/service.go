package user

type Service interface {
	// Create creates a new user
	Create(cu *CreateUser) (*User, error)

	// UserToDTO converts User model to UserDTO model
	UserToDTO(u *User) UserDTO
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(cu *CreateUser) (*User, error) {
	return s.repo.Create(cu)
}

func (s *service) UserToDTO(u *User) UserDTO {
	return UserDTO{
		ID:       u.ID,
		Username: u.Username,
	}
}
