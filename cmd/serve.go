package cmd

import (
	"github.com/shafi21064/ecom-go/config"
	"github.com/shafi21064/ecom-go/rest"
	"github.com/shafi21064/ecom-go/rest/handlers/product"
	"github.com/shafi21064/ecom-go/rest/handlers/review"
	"github.com/shafi21064/ecom-go/rest/handlers/user"
	"github.com/shafi21064/ecom-go/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()

	middleware := middleware.NewMiddlewares(cnf)
	userHandler := user.NewHandler()
	productHandler := product.NewHandler(middleware)
	reviewHandler := review.NewHandler()

	server := rest.NewServer(
		cnf,
		userHandler,
		productHandler,
		reviewHandler,
	)

	server.Start()
}
