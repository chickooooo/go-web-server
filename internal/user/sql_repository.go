package user

import "time"

type sqlRepository struct{}

func NewSQLRepository() Repository {
	return &sqlRepository{}
}

func (repo *sqlRepository) Create(cu *CreateUser) (*User, error) {
	return &User{
		ID:        1,
		Username:  "abcd",
		Password:  "abcd123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
