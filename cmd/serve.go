package cmd

import (
	"os"

	"github.com/shafi21064/ecom-go/config"
	"github.com/shafi21064/ecom-go/infra/db"
	"github.com/shafi21064/ecom-go/repo"
	"github.com/shafi21064/ecom-go/rest"
	"github.com/shafi21064/ecom-go/rest/handlers/product"
	"github.com/shafi21064/ecom-go/rest/handlers/user"
	"github.com/shafi21064/ecom-go/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewDBConnection()
	if err != nil {
		println(err)
		os.Exit(1)
	}

	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo(dbCon)

	middleware := middleware.NewMiddlewares(cnf)

	userHandler := user.NewHandler(userRepo)
	productHandler := product.NewHandler(middleware, productRepo)

	server := rest.NewServer(
		cnf,
		userHandler,
		productHandler,
	)

	server.Start()
}
