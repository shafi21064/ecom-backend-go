package cmd

import (
	"os"

	"e-com/config"
	"e-com/infra/db"
	"e-com/repo"
	"e-com/rest"
	"e-com/rest/handlers/product"
	"e-com/rest/handlers/user"
	"e-com/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewDBConnection(cnf.DBConfig)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	productRepo := repo.NewProductRepo(dbCon)
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
