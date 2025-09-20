package product

import (
	"github.com/shafi21064/ecom-go/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middlewares
}

func NewHandler(middlewares *middleware.Middlewares) *Handler {
	return &Handler{
		middlewares: middlewares,
	}
}
