package rest

import (
	"net/http"

	"github.com/shafi21064/ecom-go/rest/handlers"
	"github.com/shafi21064/ecom-go/rest/middleware"
)

func initRoutes(mux *http.ServeMux, mngr *middleware.Manager) {
	mux.Handle("GET /products", mngr.With(
		http.HandlerFunc(handlers.GetProducts),
	))

	mux.Handle("POST /products", mngr.With(
		http.HandlerFunc(handlers.CreateProduct),
	))

	mux.Handle("GET /products/{id}", mngr.With(
		http.HandlerFunc(handlers.GetProductsByID),
	))

	mux.Handle("PUT /products", mngr.With(
		http.HandlerFunc(handlers.UpdateProduct),
	))

	mux.Handle("DELETE /products/{id}", mngr.With(
		http.HandlerFunc(handlers.DeleteProduct),
	))
}
