package product

import (
	"github.com/shafi21064/ecom-go/repo"
	"github.com/shafi21064/ecom-go/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middlewares
	productRepo repo.ProductRepo
}

func NewHandler(
	middlewares *middleware.Middlewares,
	productRepo repo.ProductRepo,
	) *Handler {
	return &Handler{
		middlewares: middlewares,
		productRepo: productRepo,
	}
}
