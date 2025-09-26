package product


import "e-com/domain"

type Service interface {
	Create(product domain.Product) (*domain.Product, error)
	Get(productID int) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Update(productID int, product domain.Product) (*domain.Product, error)
	Delete(productID int) error
}
