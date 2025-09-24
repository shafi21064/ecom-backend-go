package product

import (
	"net/http"

	"github.com/shafi21064/ecom-go/rest/middleware"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, mngr *middleware.Manager) {
	mux.Handle("GET /products", mngr.With(
		http.HandlerFunc(h.GetProducts),
	))

	mux.Handle("POST /products", mngr.With(
		http.HandlerFunc(h.CreateProduct),
		h.middlewares.AuthenticateJwt,
	))

	mux.Handle("GET /products/{id}", mngr.With(
		http.HandlerFunc(h.GetProductsByID),
	))

	mux.Handle("PUT /products/{id}", mngr.With(
		http.HandlerFunc(h.UpdateProduct),
		h.middlewares.AuthenticateJwt,
	))

	mux.Handle("DELETE /products/{id}", mngr.With(
		http.HandlerFunc(h.DeleteProduct),
		h.middlewares.AuthenticateJwt,
	))
}
