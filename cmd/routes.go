package cmd

import (
	"net/http"

	"github.com/shafi21064/ecom-go/handlers"
	"github.com/shafi21064/ecom-go/middleware"
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
}
