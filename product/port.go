package product

import (
	"e-com/domain"
	productHandler "e-com/rest/handlers/product"
)

type Service interface {
	productHandler.Service
}

type ProductRepo interface {
	Create(product domain.Product) (*domain.Product, error)
	Get(productID int) (*domain.Product, error)
	List(page, limit int) ([]*domain.Product, error)
	Count() (int, error)
	Update(productID int, product domain.Product) (*domain.Product, error)
	Delete(productID int) error
}
