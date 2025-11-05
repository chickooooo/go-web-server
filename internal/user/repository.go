package user

type Repository interface {
	Create(cu *CreateUser) (*User, error)
}
