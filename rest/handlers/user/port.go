package user

import "e-com/domain"

type Service interface {
	Create(user domain.User) (*domain.User, error)
	Get(email string, password string) (*domain.User, error)
}
