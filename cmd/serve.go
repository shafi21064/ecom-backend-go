package cmd

import (
	"github.com/shafi21064/ecom-go/config"
	"github.com/shafi21064/ecom-go/rest"
	"github.com/shafi21064/ecom-go/rest/handlers/product"
	"github.com/shafi21064/ecom-go/rest/handlers/user"
)

func Serve() {
	cnf := config.GetConfig()

	userHandler := user.NewHandler()
	productHandler := product.NewHandler()

	server := rest.NewServer(userHandler, productHandler)

	server.Start(cnf)
}
